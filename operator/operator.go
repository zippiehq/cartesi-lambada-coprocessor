package operator

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ipfs_files "github.com/ipfs/boxo/files"
	ipfs_path "github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	ipfs_api "github.com/ipfs/kubo/client/rpc"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkelcontracts "github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdkmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/zippiehq/cartesi-lambada-coprocessor/aggregator"
	tm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/chainio"
	"github.com/zippiehq/cartesi-lambada-coprocessor/metrics"
)

const AVS_NAME = "lambada-coprocessor"
const SEM_VER = "0.0.1"

type Operator struct {
	config Config

	log logging.Logger

	blsKeypair   *bls.KeyPair
	ecdsaPrivKey *ecdsa.PrivateKey

	operatorAddr common.Address
	operatorId   sdktypes.OperatorId

	ethClient  eth.Client
	ipfsClient *ipfs_api.HttpApi
	// TODO(samlaf): remove both avsWriter and eigenlayerWrite from operator
	// they are only used for registration, so we should make a special registration package
	// this way, auditing this operator code makes it obvious that operators don't need to
	// write to the chain during the course of their normal operations
	// writing to the chain should be done via the cli only
	metricsReg       *prometheus.Registry
	metrics          metrics.Metrics
	nodeApi          *nodeapi.NodeApi
	avsWriter        *chainio.AvsWriter
	avsReader        chainio.AvsReaderer
	avsSubscriber    chainio.AvsSubscriberer
	eigenlayerReader sdkelcontracts.ELReader
	eigenlayerWriter sdkelcontracts.ELWriter
	// ip address of aggregator
	aggregatorServerIpPortAddr string
	// rpc client to query tasks from batch and send signed responses to aggregator
	aggregatorRpcClient AggregatorRpcClienter

	// receive new tasks in this chan (typically from listening to onchain event)
	newBatchChan chan *tm.ContractLambadaCoprocessorTaskManagerTaskBatchRegistered
}

type JSONCompute struct {
	Metadata map[string]string `json:"metadata"`
	Payload  string            `json:"payload"`
}

func NewOperator(blsPwd, ecdsaPwd string, cfg Config) (*Operator, error) {
	var logLevel logging.LogLevel
	if cfg.Production {
		logLevel = logging.Production
	} else {
		logLevel = logging.Development
	}
	logger, err := logging.NewZapLogger(logLevel)
	if err != nil {
		return nil, err
	}
	reg := prometheus.NewRegistry()
	eigenMetrics := sdkmetrics.NewEigenMetrics(AVS_NAME, cfg.EigenMetricsIpPortAddress, reg, logger)
	avsAndEigenMetrics := metrics.NewAvsAndEigenMetrics(AVS_NAME, eigenMetrics, reg)

	// Setup IPFS Api
	ipfsURL := fmt.Sprintf("http://%s", cfg.IPFSIpPortAddress)
	ipfsApi, err := ipfs_api.NewURLApiWithClient(ipfsURL, http.DefaultClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create IPFS API client - %s", err)
	}

	// Setup Node Api
	nodeApi := nodeapi.NewNodeApi(AVS_NAME, SEM_VER, cfg.NodeApiIpPortAddress, logger)

	var ethRpcClient, ethWsClient eth.Client
	if cfg.EnableMetrics {
		rpcCallsCollector := rpccalls.NewCollector(AVS_NAME, reg)
		ethRpcClient, err = eth.NewInstrumentedClient(cfg.EthRpcUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = eth.NewInstrumentedClient(cfg.EthWsUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, err
		}
	} else {
		ethRpcClient, err = eth.NewClient(cfg.EthRpcUrl)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = eth.NewClient(cfg.EthWsUrl)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, err
		}
	}

	var deployment chainio.AVSDeployment
	if err := sdkutils.ReadJsonConfig(cfg.AVSDeploymentOutputPath, &deployment); err != nil {
		return nil, fmt.Errorf("failed to read deployment parameters - %s", err)
	}

	blsKeyPair, err := bls.ReadPrivateKeyFromFile(cfg.BLSPrivateKeyStorePath, blsPwd)
	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)
		return nil, err
	}

	ecdsaKey, err := sdkecdsa.ReadKey(cfg.ECDSAPrivateKeyStorePath, ecdsaPwd)
	if err != nil {
		logger.Errorf("Cannot parse ecdsa private key", "err", err)
		return nil, err
	}
	operatorAddr := crypto.PubkeyToAddress(ecdsaKey.PublicKey)

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 cfg.EthRpcUrl,
		EthWsUrl:                   cfg.EthWsUrl,
		RegistryCoordinatorAddr:    deployment.Addresses.RegistryCoordinator,
		OperatorStateRetrieverAddr: deployment.Addresses.OperatorStateRetriever,
		AvsName:                    AVS_NAME,
		PromMetricsIpPortAddress:   cfg.EigenMetricsIpPortAddress,
	}
	operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
		cfg.ECDSAPrivateKeyStorePath,
		ecdsaPwd,
	)
	if err != nil {
		return nil, err
	}
	sdkClients, err := clients.BuildAll(chainioConfig, operatorEcdsaPrivateKey, logger)
	if err != nil {
		panic(err)
	}

	avsWriter, err := chainio.NewAvsWriter(ecdsaKey, deployment, ethRpcClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsWriter", "err", err)
		return nil, err
	}

	avsReader, err := chainio.NewAvsReader(deployment, ethRpcClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}

	avsSubscriber, err := chainio.NewAvsSubscriber(deployment, ethWsClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}

	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	quorumNames := map[sdktypes.QuorumNum]string{
		0: "quorum0",
	}
	economicMetricsCollector := economic.NewCollector(
		sdkClients.ElChainReader, sdkClients.AvsRegistryChainReader,
		AVS_NAME, logger, operatorAddr, quorumNames)
	reg.MustRegister(economicMetricsCollector)

	aggregatorRpcClient, err := NewAggregatorRpcClient(cfg.AggregatorServerIpPortAddress, logger, avsAndEigenMetrics)
	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	operator := &Operator{
		config: cfg,

		log: logger,

		blsKeypair:   blsKeyPair,
		ecdsaPrivKey: ecdsaKey,

		operatorAddr: operatorAddr,
		operatorId:   [32]byte{0}, // this is set below

		metricsReg:                 reg,
		metrics:                    avsAndEigenMetrics,
		nodeApi:                    nodeApi,
		ethClient:                  ethRpcClient,
		ipfsClient:                 ipfsApi,
		avsWriter:                  avsWriter,
		avsReader:                  avsReader,
		avsSubscriber:              avsSubscriber,
		eigenlayerReader:           sdkClients.ElChainReader,
		eigenlayerWriter:           sdkClients.ElChainWriter,
		aggregatorServerIpPortAddr: cfg.AggregatorServerIpPortAddress,
		aggregatorRpcClient:        aggregatorRpcClient,

		newBatchChan: make(chan *tm.ContractLambadaCoprocessorTaskManagerTaskBatchRegistered),
	}

	// OperatorId is set in contract during registration so we get it after registering operator.
	operatorId, err := sdkClients.AvsRegistryChainReader.GetOperatorId(&bind.CallOpts{}, operator.operatorAddr)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}
	operator.operatorId = operatorId
	logger.Info("Operator info",
		"operatorId", operatorId,
		"operatorAddr", operatorAddr,
		"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	)

	return operator, nil

}

func (o *Operator) Start(ctx context.Context) error {
	operatorIsRegistered, err := o.avsReader.IsOperatorRegistered(&bind.CallOpts{}, o.operatorAddr)
	if err != nil {
		o.log.Error("Error checking if operator is registered", "err", err)
		return err
	}
	if !operatorIsRegistered {
		// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack trace
		// that hides the actual error message. This error msg is more explicit and doesn't require showing a stack trace to the user.
		return fmt.Errorf("operator is not registered. Registering operator using the operator-cli before starting operator")
	}

	o.log.Infof("Starting operator.")

	if o.config.EnableNodeApi {
		o.nodeApi.Start()
	}
	var metricsErrChan <-chan error
	if o.config.EnableMetrics {
		metricsErrChan = o.metrics.Start(ctx, o.metricsReg)
	} else {
		metricsErrChan = make(chan error, 1)
	}

	// TODO(samlaf): wrap this call with increase in avs-node-spec metric
	sub, err := o.avsSubscriber.SubscribeToNewBatches(o.newBatchChan)
	if err != nil {
		o.log.Errorf("failed to subscribe to task batches - %s", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-metricsErrChan:
			// TODO(samlaf); we should also register the service as unhealthy in the node api
			// https://eigen.nethermind.io/docs/spec/api/
			o.log.Fatal("Error in metrics server", "err", err)
		case err := <-sub.Err():
			o.log.Error("Error in websocket subscription", "err", err)
			// TODO(samlaf): write unit tests to check if this fixed the issues we were seeing
			sub.Unsubscribe()
			// TODO(samlaf): wrap this call with increase in avs-node-spec metric
			sub, err = o.avsSubscriber.SubscribeToNewBatches(o.newBatchChan)
			if err != nil {
				o.log.Errorf("failed to subscribe to task batches - %s", err)
			}
		case newTaskCreatedLog := <-o.newBatchChan:
			o.metrics.IncNumTasksReceived()
			if err := o.processTaskBatch(newTaskCreatedLog); err != nil {
				o.log.Errorf("failed to process task batch - %s", err)
			}
		}
	}
}

// Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (o *Operator) processTaskBatch(newBatch *tm.ContractLambadaCoprocessorTaskManagerTaskBatchRegistered) error {
	o.log.Debug("Received new task batch", "task", newBatch)

	o.log.Info("Received new task batch",
		"index", newBatch.Batch.Index,
		"blockNumber", newBatch.Batch.BlockNumber,
		"merkleRoot", newBatch.Batch.MerkeRoot,
		"quorumNumbers", newBatch.Batch.QuorumNumbers,
		`quroumThresholdPrecentage`, newBatch.Batch.QuorumThresholdPercentage,
	)

	tasks, err := o.aggregatorRpcClient.GetBatchTasks(newBatch.Batch.Index)
	if err != nil {
		return err
	}

	for _, t := range tasks {
		cid, output, err := o.computeTaskOutput(t, newBatch.Batch.BlockNumber)
		if err != nil {
			cid, output = fakeOutput(len(t.Input))
			o.log.Errorf("failed to request echo from lambada service -%s, faking result - %s, %s",
				err, cid, string(output))
		}

		if err = o.sendTaskOutput(t.Index, cid, hashOutput(output)); err != nil {
			return err
		}
	}

	return nil
}

func (o *Operator) computeTaskOutput(t core.Task, blockNumber uint32) (string, []byte, error) {
	// Query lambada compute endpoint.
	taskCID := string(t.ProgramID)
	requestURL := fmt.Sprintf("http://%s/compute/%s?json=true",
		o.config.LambadaIpPortAddress,
		taskCID,
	)

	blockNumberBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(blockNumberBytes, uint64(blockNumber))

	curBlock, err := o.ethClient.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		o.log.Errorf("Unable to get current block")
		return "", nil, err
	}

	base64Hash := func(d []byte) string {
		hash := sha256.Sum256(d)
		return base64.StdEncoding.EncodeToString(hash[:])
	}

	jsonCompute := JSONCompute{
		Metadata: map[string]string{
			base64Hash([]byte("sequencer")):                      base64Hash([]byte("coprocessor")),
			base64Hash([]byte("coprocessor-batch-block-number")): base64.StdEncoding.EncodeToString(blockNumberBytes),
			base64Hash([]byte("coprocessor-batch-block-hash")):   base64.StdEncoding.EncodeToString(curBlock.Hash().Bytes()),
		},
		Payload: base64.StdEncoding.EncodeToString(t.Input),
	}

	input, err := json.Marshal(jsonCompute)
	if err != nil {
		return "", nil, err
	}

	o.log.Infof("sending request to lambada instance - %s", requestURL)

	resp, err := http.Post(requestURL, "application/octet-stream", bytes.NewBuffer(input))
	if err != nil {
		return "", nil, err
	}

	// Parse CID with echo output.
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}
	jsonCID := struct {
		CID string `json:"cid"`
	}{}
	if err = json.Unmarshal(respData, &jsonCID); err != nil {
		return "", nil, err
	}
	cid, err := cid.Decode(jsonCID.CID)
	if err != nil {
		return "", nil, err
	}

	// Query echo output from IPFS.
	outputPath, err := ipfs_path.NewPath(fmt.Sprintf("/ipfs/%s/output", cid.String()))
	if err != nil {
		return cid.String(), []byte{}, nil
	}
	outputNode, err := o.ipfsClient.Unixfs().Get(context.TODO(), outputPath)
	if err != nil {
		return cid.String(), []byte{}, nil
	}
	outputFile := ipfs_files.ToFile(outputNode)
	defer outputFile.Close()
	output, err := io.ReadAll(outputFile)
	if err != nil {
		return "", nil, err
	}

	return cid.String(), output, nil
}

func (o *Operator) sendTaskOutput(taskIdx sdktypes.TaskIndex, resultCID string, outputHash [32]byte) error {
	// Sign task response;
	cid, err := cid.Decode(resultCID)
	if err != nil {
		return fmt.Errorf("failed to decode result CID - %s", err)
	}
	resp := tm.ILambadaCoprocessorTaskManagerTaskResponse{
		ResultCID:  cid.Bytes(),
		OutputHash: outputHash,
	}

	taskResponseHash, err := core.GetTaskResponseDigest(&resp)
	if err != nil {
		return fmt.Errorf("failed to compute task response digest - %s", err)
	}

	blsSignature := o.blsKeypair.SignMessage(taskResponseHash)

	// Send task to aggregator
	signedTaskResponse := &aggregator.SignedTaskResponse{
		TaskResponse: core.TaskResponse{
			TaskIndex:  taskIdx,
			OperatorID: o.operatorId,
			OutputHash: resp.OutputHash[:],
			ResultCID:  resp.ResultCID,
		},
		BlsSignature: *blsSignature,
	}
	go o.aggregatorRpcClient.SendSignedTaskResponseToAggregator(signedTaskResponse)

	return nil
}

func fakeOutput(size int) (string, []byte) {
	cid := "bafybeif5liyov4oltxvc34qh3uyyxy5znisbmarsorxrp4nr3o7ywjuo5u"
	output := make([]byte, size)
	rand.Read(output)
	return cid, output
}

func hashOutput(output []byte) [32]byte {
	outputHash := sha256.Sum256(output)
	return outputHash
}

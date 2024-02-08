package operator

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-faster/xor"
	ipfs_files "github.com/ipfs/boxo/files"
	ipfs_path "github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	ipfs_api "github.com/ipfs/kubo/client/rpc"
	"github.com/prometheus/client_golang/prometheus"

	sdkavsregistry "github.com/Layr-Labs/eigensdk-go/chainio/avsregistry"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdkelcontracts "github.com/Layr-Labs/eigensdk-go/chainio/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdkmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/Layr-Labs/eigensdk-go/signer"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/zippiehq/cartesi-lambada-coprocessor/aggregator"
	tm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/chainio"
	"github.com/zippiehq/cartesi-lambada-coprocessor/metrics"
	"github.com/zippiehq/cartesi-lambada-coprocessor/types"
)

const AVS_NAME = "lambada-coprocessor"
const SEM_VER = "0.0.1"

type Operator struct {
	config types.NodeConfig

	operatorAddr common.Address
	operatorId   bls.OperatorId
	blsKeypair   *bls.KeyPair

	// needed when opting in to avs (allow this service manager contract to slash operator)
	serviceManagerAddr common.Address

	log        logging.Logger
	ethClient  eth.EthClient
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

func NewOperatorFromConfig(c types.NodeConfig) (*Operator, error) {
	var logLevel logging.LogLevel
	if c.Production {
		logLevel = logging.Production
	} else {
		logLevel = logging.Development
	}
	logger, err := logging.NewZapLogger(logLevel)
	if err != nil {
		return nil, err
	}
	reg := prometheus.NewRegistry()
	eigenMetrics := sdkmetrics.NewEigenMetrics(AVS_NAME, c.EigenMetricsIpPortAddress, reg, logger)
	avsAndEigenMetrics := metrics.NewAvsAndEigenMetrics(AVS_NAME, eigenMetrics, reg)

	// Setup IPFS Api
	ipfsURL := fmt.Sprintf("http://%s", os.Getenv("IPFS_ADDRESS"))
	ipfsApi, err := ipfs_api.NewURLApiWithClient(ipfsURL, http.DefaultClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create IPFS API client - %s", err)
	}

	// Setup Node Api
	nodeApi := nodeapi.NewNodeApi(AVS_NAME, SEM_VER, c.NodeApiIpPortAddress, logger)

	var ethRpcClient, ethWsClient eth.EthClient
	if c.EnableMetrics {
		rpcCallsCollector := rpccalls.NewCollector(AVS_NAME, reg)
		ethRpcClient, err = eth.NewInstrumentedClient(c.EthRpcUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = eth.NewInstrumentedClient(c.EthWsUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, err
		}
	} else {
		ethRpcClient, err = eth.NewClient(c.EthRpcUrl)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = eth.NewClient(c.EthWsUrl)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, err
		}
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)
		return nil, err
	}
	// TODO(samlaf): should we add the chainId to the config instead?
	// this way we can prevent creating a signer that signs on mainnet by mistake
	// if the config says chainId=5, then we can only create a goerli signer
	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	sgn, err := signer.NewPrivateKeyFromKeystoreSigner(c.EcdsaPrivateKeyStorePath, ecdsaKeyPassword, chainId)
	if err != nil {
		logger.Errorf("Cannot create signer", "err", err)
		return nil, err
	}

	avsWriter, err := chainio.NewAvsWriter(sgn, common.HexToAddress(c.AVSServiceManagerAddress),
		common.HexToAddress(c.BLSOperatorStateRetrieverAddress), ethRpcClient, logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsWriter", "err", err)
		return nil, err
	}

	avsServiceBindings, err := chainio.NewAvsServiceBindings(
		common.HexToAddress(c.AVSServiceManagerAddress),
		common.HexToAddress(c.BLSOperatorStateRetrieverAddress),
		ethRpcClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	blsRegistryCoordinatorAddr, err := avsServiceBindings.ServiceManager.RegistryCoordinator(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	stakeRegistryAddr, err := avsServiceBindings.ServiceManager.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	blsPubkeyRegistryAddr, err := avsServiceBindings.ServiceManager.BlsPubkeyRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	avsRegistryContractClient, err := sdkclients.NewAvsRegistryContractsChainClient(
		blsRegistryCoordinatorAddr, common.HexToAddress(c.BLSOperatorStateRetrieverAddress), stakeRegistryAddr, blsPubkeyRegistryAddr, ethRpcClient, logger,
	)
	if err != nil {
		return nil, err
	}
	avsRegistryReader, err := sdkavsregistry.NewAvsRegistryReader(avsRegistryContractClient, logger, ethRpcClient)
	if err != nil {
		return nil, err
	}
	avsReader, err := chainio.NewAvsReader(avsRegistryReader, avsServiceBindings, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}
	avsSubscriber, err := chainio.NewAvsSubscriber(common.HexToAddress(c.AVSServiceManagerAddress),
		common.HexToAddress(c.BLSOperatorStateRetrieverAddress), ethWsClient, logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}

	slasherAddr, err := avsReader.AvsServiceBindings.ServiceManager.Slasher(&bind.CallOpts{})
	if err != nil {
		logger.Error("Cannot get slasher address", "err", err)
		return nil, err
	}

	elContractsClient, err := sdkclients.NewELContractsChainClient(slasherAddr, common.HexToAddress(c.BlsPublicKeyCompendiumAddress), ethRpcClient, ethWsClient, logger)
	if err != nil {
		logger.Error("Cannot create ELContractsChainClient", "err", err)
		return nil, err
	}

	eigenlayerWriter := sdkelcontracts.NewELChainWriter(elContractsClient, ethRpcClient, sgn, logger, eigenMetrics)
	if err != nil {
		logger.Error("Cannot create EigenLayerWriter", "err", err)
		return nil, err
	}
	eigenlayerReader, err := sdkelcontracts.NewELChainReader(elContractsClient, logger, ethRpcClient)
	if err != nil {
		logger.Error("Cannot create EigenLayerReader", "err", err)
		return nil, err
	}

	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	quorumNames := map[sdktypes.QuorumNum]string{
		0: "quorum0",
	}
	economicMetricsCollector := economic.NewCollector(eigenlayerReader, avsRegistryReader, AVS_NAME, logger, common.HexToAddress(c.OperatorAddress), quorumNames)
	reg.MustRegister(economicMetricsCollector)

	aggregatorRpcClient, err := NewAggregatorRpcClient(c.AggregatorServerIpPortAddress, logger, avsAndEigenMetrics)
	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	operator := &Operator{
		config: c,

		operatorAddr: common.HexToAddress(c.OperatorAddress),
		operatorId:   [32]byte{0}, // this is set below
		blsKeypair:   blsKeyPair,

		serviceManagerAddr: common.HexToAddress(c.AVSServiceManagerAddress),

		log:                        logger,
		metricsReg:                 reg,
		metrics:                    avsAndEigenMetrics,
		nodeApi:                    nodeApi,
		ethClient:                  ethRpcClient,
		ipfsClient:                 ipfsApi,
		avsWriter:                  avsWriter,
		avsReader:                  avsReader,
		avsSubscriber:              avsSubscriber,
		eigenlayerReader:           eigenlayerReader,
		eigenlayerWriter:           eigenlayerWriter,
		aggregatorServerIpPortAddr: c.AggregatorServerIpPortAddress,
		aggregatorRpcClient:        aggregatorRpcClient,

		newBatchChan: make(chan *tm.ContractLambadaCoprocessorTaskManagerTaskBatchRegistered),
	}

	if c.RegisterOperatorOnStartup {
		operator.registerOperatorOnStartup(common.HexToAddress(c.BlsPublicKeyCompendiumAddress))
	}

	// OperatorId is set in contract during registration so we get it after registering operator.
	operatorId, err := avsRegistryReader.GetOperatorId(context.Background(), operator.operatorAddr)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}
	operator.operatorId = operatorId
	logger.Info("Operator info",
		"operatorId", operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	)

	return operator, nil

}

func (o *Operator) Start(ctx context.Context) error {
	operatorIsRegistered, err := o.avsReader.IsOperatorRegistered(ctx, o.operatorAddr)
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
		cid, output, err := o.computeTaskOutput(t.Input)
		if err != nil {
			cid, output = fakeOutput(len(t.Input))
			o.log.Errorf("failed to request echo from lambada service -%s, faking result - %s, %s",
				err, cid, string(output))
		}

		if err = o.sendTaskOutput(t.Index, hashOutput(cid, output)); err != nil {
			return err
		}
	}

	return nil
}

func (o *Operator) computeTaskOutput(input []byte) (string, []byte, error) {
	// Query lambada compute endpoint.
	requestURL := fmt.Sprintf("http://%s/compute/%s",
		os.Getenv("LAMBADA_ADDRESS"),
		os.Getenv("LAMBADA_COMPUTE_CID"),
	)
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
		return "", nil, err
	}
	outputNode, err := o.ipfsClient.Unixfs().Get(context.TODO(), outputPath)
	if err != nil {
		return "", nil, err
	}
	outputFile := ipfs_files.ToFile(outputNode)
	defer outputFile.Close()
	output, err := io.ReadAll(outputFile)
	if err != nil {
		return "", nil, err
	}

	return cid.String(), output, nil
}

func (o *Operator) sendTaskOutput(taskIdx sdktypes.TaskIndex, outputHash [32]byte) error {
	// Sign task response;
	resp := tm.ILambadaCoprocessorTaskManagerTaskResponse{
		OutputHash: outputHash,
	}
	taskResponseHash, err := core.GetTaskResponseDigest(&resp)
	if err != nil {
		return fmt.Errorf("failed to compute task response digest - %s", err)
	}
	blsSignature := o.blsKeypair.SignMessage(taskResponseHash)

	// Send task to aggregator
	signedTaskResponse := &aggregator.SignedTaskResponse{
		ILambadaCoprocessorTaskManagerTaskResponse: resp,
		TaskIndex:    taskIdx,
		BlsSignature: *blsSignature,
		OperatorId:   o.operatorId,
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

func hashOutput(ouputCID string, output []byte) [32]byte {
	cidHash := sha256.Sum256([]byte(ouputCID))
	outputHash := sha256.Sum256(output)
	var hash [32]byte
	xor.Bytes(hash[:], cidHash[:], outputHash[:])
	return hash
}

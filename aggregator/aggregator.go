package aggregator

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"

	tm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/chainio"
)

const (
	// all operators in quorum0 must sign the task response in order for it to be accepted
	// TODO: our contracts require uint8 but right now sdktypes.QuorumThresholdPercentage is uint8 prob need to update our inc-sq contracts to use uint8 as well?
	QUORUM_THRESHOLD_NUMERATOR   = sdktypes.QuorumThresholdPercentage(100)
	QUORUM_THRESHOLD_DENOMINATOR = sdktypes.QuorumThresholdPercentage(100)

	// task batch parameters
	BATCH_PERIOD = 10 * time.Second

	// number of blocks after which a task is considered expired
	// this hardcoded here because it's also hardcoded in the contracts, but should
	// ideally be fetched from the contracts
	TASK_CHALLENGE_WIDNOW_BLOCK = 100
	BLOCK_TIME_SECONDS          = 12 * time.Second

	AVS_NAME = "lambada-coprocessor"
)

// we only use a single quorum (quorum 0) for incredible squaring
var QUORUM_NUMBERS = sdktypes.QuorumNums{0}

// Aggregator sends tasks (input to echo) onchain, then listens for operator signed TaskResponses.
// It aggregates responses signatures, and if any of the TaskResponses reaches the QuorumThresholdPercentage for each quorum
// (currently we only use a single quorum of the ERC20Mock token), it sends the aggregated TaskResponse and signature onchain.
//
// The signature is checked in the BLSSignatureChecker.sol contract, which expects a
//
//	struct NonSignerStakesAndSignature {
//		uint32[] nonSignerQuorumBitmapIndices;
//		BN254.G1Point[] nonSignerPubkeys;
//		BN254.G1Point[] quorumApks;
//		BN254.G2Point apkG2;
//		BN254.G1Point sigma;
//		uint32[] quorumApkIndices;
//		uint32[] totalStakeIndices;
//		uint32[][] nonSignerStakeIndices; // nonSignerStakeIndices[quorumNumberIndex][nonSignerIndex]
//	}
//
// A task can only be responded onchain by having enough operators sign on it such that their stake in each quorum reaches the QuorumThresholdPercentage.
// In order to verify this onchain, the Registry contracts store the history of stakes and aggregate pubkeys (apks) for each operators and each quorum. These are
// updated everytime an operator registers or deregisters with the BLSRegistryCoordinatorWithIndices.sol contract, or calls UpdateStakes() on the StakeRegistry.sol contract,
// after having received new delegated shares or having delegated shares removed by stakers queuing withdrawals. Each of these pushes to their respective datatype array a new entry.
//
// This is true for quorumBitmaps (represent the quorums each operator is opted into), quorumApks (apks per quorum), totalStakes (total stake per quorum), and nonSignerStakes (stake per quorum per operator).
// The 4 "indices" in NonSignerStakesAndSignature basically represent the index at which to fetch their respective data, given a blockNumber at which the task was created.
// Note that different data types might have different indices, since for eg QuorumBitmaps are updated for operators registering/deregistering, but not for UpdateStakes.
// Thankfully, we have deployed a helper contract BLSOperatorStateRetriever.sol whose function getCheckSignaturesIndices() can be used to fetch the indices given a block number.
//
// The 4 other fields nonSignerPubkeys, quorumApks, apkG2, and sigma, however, must be computed individually.
// apkG2 and sigma are just the aggregated signature and pubkeys of the operators who signed the task response (aggregated over all quorums, so individual signatures might be duplicated).
// quorumApks are the G1 aggregated pubkeys of the operators who signed the task response, but one per quorum, as opposed to apkG2 which is summed over all quorums.
// nonSignerPubkeys are the G1 pubkeys of the operators who did not sign the task response, but were opted into the quorum at the blocknumber at which the task was created.
// Upon sending a task onchain (or receiving a NewTaskCreated Event if the tasks were sent by an external task generator), the aggregator can get the list of all operators opted into each quorum at that
// block number by calling the getOperatorState() function of the BLSOperatorStateRetriever.sol contract.
type Aggregator struct {
	log logging.Logger

	apiServerAddr string

	storage   Storage
	avsWriter chainio.AvsWriterer
	blsAgg    blsagg.BlsAggregationService
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregator(privKey, dbPwd string, cfg Config, log logging.Logger) (*Aggregator, error) {
	storageCfg := MySqlStorageConfig{
		Address:  cfg.DatabaseAddress,
		Database: cfg.DatabaseName,
		User:     cfg.DatabaseUser,
		Password: dbPwd,
	}

	storage, err := NewMySqlStorage(storageCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage - %s", err)
	}

	var deployment chainio.AVSDeployment
	if err := sdkutils.ReadJsonConfig(cfg.AVSDeploymentOutputPath, &deployment); err != nil {
		return nil, fmt.Errorf("failed to read AVS deployment file - %s", err)
	}

	if privKey[:2] == "0x" {
		privKey = privKey[2:]
	}
	ecdsaPrivKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ECDSA private key - %s", err)
	}

	chainioConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 cfg.EthHttpRpcUrl,
		EthWsUrl:                   cfg.EthWsRpcUrl,
		RegistryCoordinatorAddr:    deployment.Addresses.RegistryCoordinator,
		OperatorStateRetrieverAddr: deployment.Addresses.OperatorStateRetriever,
		AvsName:                    AVS_NAME,
		PromMetricsIpPortAddress:   ":9090",
	}
	clients, err := sdkclients.BuildAll(chainioConfig, ecdsaPrivKey, log)
	if err != nil {
		return nil, fmt.Errorf("failed to create sdk clients - %s", err)
	}

	ethClient, err := eth.NewClient(cfg.EthHttpRpcUrl)
	if err != nil {
		return nil, fmt.Errorf("faield to create ETH RPC client - %s", err)
	}
	avsReader, err := chainio.NewAvsReader(deployment, ethClient, log)
	if err != nil {
		return nil, fmt.Errorf("failed to create AVS reader - %s", err)
	}
	avsWriter, err := chainio.NewAvsWriter(ecdsaPrivKey, deployment, ethClient, log)
	if err != nil {
		return nil, fmt.Errorf("failed to create AVS writer - %s", err)
	}

	operatorPubkeysService := oprsinfoserv.NewOperatorsInfoServiceInMemory(
		context.Background(),
		clients.AvsRegistryChainSubscriber,
		clients.AvsRegistryChainReader,
		log,
	)

	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(
		avsReader,
		operatorPubkeysService,
		log,
	)

	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, log)

	return &Aggregator{
		log:           log,
		apiServerAddr: cfg.APIServerAddress,

		storage:   storage,
		avsWriter: avsWriter,
		blsAgg:    blsAggregationService,
	}, nil
}

func (agg *Aggregator) Start(ctx context.Context) error {
	agg.log.Infof("Starting aggregator.")
	agg.log.Infof("Starting aggregator rpc server.")
	go agg.startServer(ctx)
	go agg.startAPIServer()

	batchTick := time.NewTicker(BATCH_PERIOD)
	defer batchTick.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-batchTick.C:
			if err := agg.createTaskBatch(); err != nil {
				agg.log.Errorf("failed to create task batch - %s", err)
			}
		case blsAggServiceResp := <-agg.blsAgg.GetResponseChannel():
			agg.log.Info("Received response from blsAggregationService", "blsAggServiceResp", blsAggServiceResp)
			if err := agg.sendAggregatedResponseToContract(blsAggServiceResp); err != nil {
				agg.log.Errorf("failed to send aggregated response to contract - %s", err)
			}
		}
	}
}

func (agg *Aggregator) addTask(programID []byte, input []byte) (sdktypes.TaskIndex, error) {
	t := core.Task{
		ProgramID: programID,
		Input:     input,
		InputHash: core.Keccack256(input),
	}
	return agg.storage.AddPendingTask(t)
}

func (agg *Aggregator) createTaskBatch() error {
	// Create new batch from pedning tasks.
	pendingTasks, err := agg.storage.AllPendingTasks()
	if err != nil {
		return err
	}
	if len(pendingTasks) == 0 {
		return nil
	}
	batchTasks, merkle, err := core.BuildTaskBatchMerkle(pendingTasks)
	if err != nil {
		return err
	}

	// Post new batch onchain.
	onchainBatch, err := agg.avsWriter.RegisterNewTaskBatch(
		context.TODO(), [32]byte(merkle.GetRoot()), QUORUM_THRESHOLD_NUMERATOR, QUORUM_NUMBERS,
	)
	if err != nil {
		return err
	}

	// Write new batch to storage.
	batch := core.TaskBatch{
		Index:                     onchainBatch.Index,
		BlockNumber:               onchainBatch.BlockNumber,
		QuorumNumbers:             onchainBatch.QuorumNumbers,
		QuorumThresholdPercentage: onchainBatch.QuorumThresholdPercentage,
		Tasks:                     batchTasks,
		Merkle:                    merkle,
	}
	if err := agg.storage.AddTaskBatch(batch); err != nil {
		return err
	}

	// Start accepting singed responses for every task in batch.
	quorumThresholdPercentages := make(sdktypes.QuorumThresholdPercentages, len(batch.QuorumNumbers))
	for i, _ := range batch.QuorumNumbers {
		quorumThresholdPercentages[i] = sdktypes.QuorumThresholdPercentage(batch.QuorumThresholdPercentage)
	}
	// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
	// and it should monitor the chain and only expire the task aggregation once the chain has reached that block number.
	taskTimeToExpiry := TASK_CHALLENGE_WIDNOW_BLOCK * BLOCK_TIME_SECONDS
	for _, t := range batch.Tasks {
		var quorumNums sdktypes.QuorumNums
		for _, quorumNum := range batch.QuorumNumbers {
			quorumNums = append(quorumNums, sdktypes.QuorumNum(quorumNum))
		}
		if err := agg.blsAgg.InitializeNewTask(
			t.Index, batch.BlockNumber, quorumNums, quorumThresholdPercentages, taskTimeToExpiry,
		); err != nil {
			agg.log.Errorf("failed to initialize bls aggregation for task %d from batch %d", t.Index, batch.Index)
		}
	}

	return nil
}

func (agg *Aggregator) getBatchTasks(batchIdx core.TaskBatchIndex) ([]core.Task, error) {
	b, err := agg.storage.TaskBatch(batchIdx)
	if err != nil {
		return nil, err
	}

	return b.Tasks, nil
}

func (agg *Aggregator) processTaskResponse(
	bi core.TaskBatchIndex,
	t core.Task,
	r core.TaskResponse,
	s bls.Signature,
) error {
	if err := agg.storage.AddTaskResponse(r); err != nil {
		return err
	}

	responseDigest, err := core.TaskResponseSigHash(bi, t, r)
	if err != nil {
		return err
	}
	return agg.blsAgg.ProcessNewSignature(context.Background(), r.TaskIndex, responseDigest, &s, r.OperatorID)
}

func (agg *Aggregator) sendAggregatedResponseToContract(
	resp blsagg.BlsAggregationServiceResponse,
) error {
	// TODO: check if blsAggServiceResp contains an err
	if resp.Err != nil {
		agg.log.Error("BlsAggregationServiceResponse contains an error", "err", resp.Err)
		// panicing to help with debugging (fail fast), but we shouldn't panic if we run this in production
		panic(resp.Err)
	}
	nonSignerPubkeys := []tm.BN254G1Point{}
	for _, nonSignerPubkey := range resp.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, core.ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []tm.BN254G1Point{}
	for _, quorumApk := range resp.QuorumApksG1 {
		quorumApks = append(quorumApks, core.ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := tm.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(resp.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(resp.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: resp.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             resp.QuorumApkIndices,
		TotalStakeIndices:            resp.TotalStakeIndices,
		NonSignerStakeIndices:        resp.NonSignerStakeIndices,
	}

	agg.log.Info("Threshold reached. Sending aggregated response onchain.",
		"taskIndex", resp.TaskIndex,
	)

	// Ensure that sig aggregation result is consistent with storage.
	task, batchIdx, err := agg.storage.SubmittedTask(resp.TaskIndex)
	if err != nil {
		panic("submitted task not found in storage")
	}
	batch, err := agg.storage.TaskBatch(batchIdx)
	if err != nil {
		panic("task batch not found in storage")
	}
	taskResponses, err := agg.storage.TaskResponses(resp.TaskIndex)
	if err != nil {
		return err
	}
	var (
		taskResp      core.TaskResponse
		taskRespFound bool
	)
	for _, tr := range taskResponses {
		digest, err := core.TaskResponseSigHash(batch.Index, task, tr)
		if err != nil {
			return err
		}
		if digest == resp.TaskResponseDigest {
			taskResp = tr
			taskRespFound = true
		}
	}
	if !taskRespFound {
		panic("task response not found in storage")
	}

	// Generate proof for task.
	taskProof, err := core.TaskProof(batch, resp.TaskIndex)
	if err != nil {
		return err
	}
	taskProofOnchain := make([][32]byte, len(taskProof))
	for i, p := range taskProof {
		taskProofOnchain[i] = [32]byte(p)
	}

	// Post task response onchain.
	_, err = agg.avsWriter.RespondTask(
		context.Background(),
		tm.ILambadaCoprocessorTaskManagerTaskBatch{
			Index:                     batch.Index,
			BlockNumber:               batch.BlockNumber,
			MerkeRoot:                 [32]byte(batch.Merkle.GetRoot()),
			QuorumNumbers:             batch.QuorumNumbers,
			QuorumThresholdPercentage: batch.QuorumThresholdPercentage,
		},
		tm.ILambadaCoprocessorTaskManagerTask{
			BatchIndex: batch.Index,
			ProgramId:  task.ProgramID,
			InputHash:  task.InputHash,
		},
		taskProofOnchain,
		tm.ILambadaCoprocessorTaskManagerTaskResponse{
			OutputHash: [32]byte(taskResp.OutputHash),
			ResultCID:  taskResp.ResultCID,
		},
		nonSignerStakesAndSignature,
	)

	return err
}

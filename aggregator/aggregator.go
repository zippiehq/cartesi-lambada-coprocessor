package aggregator

import (
	"cmp"
	"context"
	"slices"
	"sync"
	"time"

	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signer"

	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkelcontracts "github.com/Layr-Labs/eigensdk-go/chainio/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	pubkeycompserv "github.com/Layr-Labs/eigensdk-go/services/pubkeycompendium"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/zippiehq/cartesi-lambada-coprocessor/aggregator/types"
	tm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/chainio"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/config"
)

const (
	// task batch parameters
	batchPeriod = 10 * time.Second

	// number of blocks after which a task is considered expired
	// this hardcoded here because it's also hardcoded in the contracts, but should
	// ideally be fetched from the contracts
	taskChallengeWindowBlock = 100
	blockTimeSeconds         = 12 * time.Second
)

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

	serverIpPortAddr string

	avsWriter             chainio.AvsWriterer
	blsAggregationService blsagg.BlsAggregationService

	taskMu sync.RWMutex

	taskIndex      sdktypes.TaskIndex
	pendingTasks   map[sdktypes.TaskIndex]tm.ILambadaCoprocessorTaskManagerTask
	submittedTasks map[sdktypes.TaskIndex]tm.ILambadaCoprocessorTaskManagerTask

	batchIndex types.TaskBatchIndex
	batches    map[types.TaskBatchIndex]types.TaskBatch

	taskResponses map[sdktypes.TaskIndex]map[sdktypes.TaskResponseDigest]tm.ILambadaCoprocessorTaskManagerTaskResponse
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregator(c *config.Config) (*Aggregator, error) {

	avsReader, err := chainio.NewAvsReaderFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create EthReader", "err", err)
		return nil, err
	}

	chainId, err := c.EthHttpClient.ChainID(context.Background())
	if err != nil {
		c.Logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	privateKeySigner, err := signer.NewPrivateKeySigner(c.EcdsaPrivateKey, chainId)
	if err != nil {
		c.Logger.Error("Cannot create signer", "err", err)
		return nil, err
	}
	c.Signer = privateKeySigner

	avsWriter, err := chainio.NewAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create EthSubscriber", "err", err)
		return nil, err
	}

	c.Logger.Info("SlasherAddr", "SlasherAddr", c.SlasherAddr)

	// TODO (Soubhik): This is a hack to get the slasher address. We should be able to get it from the config.
	slasherAddr, err := avsWriter.AvsContractBindings.ServiceManager.Slasher(&bind.CallOpts{})
	if err != nil {
		c.Logger.Error("Cannot get slasher address", "err", err)
		return nil, err
	}
	c.Logger.Info("BlsPublicKeyCompendiumAddress", "BlsPublicKeyCompendiumAddress", c.BlsPublicKeyCompendiumAddress)

	elContractsClient, err := sdkclients.NewELContractsChainClient(slasherAddr, c.BlsPublicKeyCompendiumAddress, c.EthHttpClient, c.EthWsClient, c.Logger)
	if err != nil {
		c.Logger.Error("Cannot create ELContractsChainClient", "err", err)
		return nil, err
	}
	eigenlayerReader, err := sdkelcontracts.NewELChainReader(elContractsClient, c.Logger, c.EthHttpClient)
	if err != nil {
		c.Logger.Error("Cannot create EigenlayerReader", "err", err)
		return nil, err
	}
	eigenlayerSubscriber, err := sdkelcontracts.NewELChainSubscriber(
		elContractsClient,
		c.Logger,
	)
	if err != nil {
		c.Logger.Error("Cannot create EigenlayerEthSubscriber", "err", err)
		return nil, err
	}

	pubkeyCompendiumService := pubkeycompserv.NewPubkeyCompendiumInMemory(context.Background(), eigenlayerSubscriber, eigenlayerReader, c.Logger)
	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(avsReader, eigenlayerReader, pubkeyCompendiumService, c.Logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, c.Logger)

	return &Aggregator{
		log:              c.Logger,
		serverIpPortAddr: c.AggregatorServerIpPortAddr,

		avsWriter:             avsWriter,
		blsAggregationService: blsAggregationService,

		taskIndex:      0,
		pendingTasks:   make(map[sdktypes.TaskIndex]tm.ILambadaCoprocessorTaskManagerTask),
		submittedTasks: make(map[sdktypes.TaskIndex]tm.ILambadaCoprocessorTaskManagerTask),

		batchIndex: 0,
		batches:    make(map[types.TaskBatchIndex]types.TaskBatch),

		taskResponses: make(map[sdktypes.TaskIndex]map[sdktypes.TaskResponseDigest]tm.ILambadaCoprocessorTaskManagerTaskResponse),
	}, nil
}

func (agg *Aggregator) Start(ctx context.Context) error {
	agg.log.Infof("Starting aggregator.")
	agg.log.Infof("Starting aggregator rpc server.")
	go agg.startServer(ctx)
	go agg.startAPIServer()

	batchTick := time.NewTicker(10 * time.Second)
	defer batchTick.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-batchTick.C:
			if err := agg.createTaskBatch(); err != nil {
				agg.log.Errorf("failed to create task batch - %s", err)
			}
		case blsAggServiceResp := <-agg.blsAggregationService.GetResponseChannel():
			agg.log.Info("Received response from blsAggregationService", "blsAggServiceResp", blsAggServiceResp)
			if err := agg.sendAggregatedResponseToContract(blsAggServiceResp); err != nil {
				agg.log.Errorf("failed to send aggregated response to contract - %s", err)
			}
		}
	}
}

func (agg *Aggregator) addTask(input []byte) (sdktypes.TaskIndex, error) {
	agg.taskMu.Lock()
	defer agg.taskMu.Unlock()

	agg.pendingTasks[agg.taskIndex] = tm.ILambadaCoprocessorTaskManagerTask{
		Input: input,
	}
	agg.taskIndex++

	return agg.taskIndex, nil
}

func (agg *Aggregator) createTaskBatch() error {
	batchTasks, batchMerkle, err := agg.makeBatch()
	if err != nil {
		return err
	}

	// Post new batch onchain.
	onchainBatch, err := agg.avsWriter.RegisterNewTaskBatch(
		context.TODO(), [32]byte(batchMerkle.GetRoot()), types.QUORUM_THRESHOLD_NUMERATOR, types.QUORUM_NUMBERS,
	)
	if err != nil {
		return err
	}

	batch := agg.confirmBatch(onchainBatch, batchTasks, batchMerkle)

	// Start accepting singed responses for every task in batch.
	quorumThresholdPercentages := make([]uint32, len(batch.QuorumNumbers))
	for i, _ := range batch.QuorumNumbers {
		quorumThresholdPercentages[i] = batch.QuorumThresholdPercentage
	}
	// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
	// and it should monitor the chain and only expire the task aggregation once the chain has reached that block number.
	taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds
	for _, t := range batch.Tasks {
		if err := agg.blsAggregationService.InitializeNewTask(
			t.Index, batch.BlockNumber, batch.QuorumNumbers, quorumThresholdPercentages, taskTimeToExpiry,
		); err != nil {
			agg.log.Errorf("failed to initialize bls aggregation for task %d from batch %d", t.Index, batch.Index)
		}
	}

	return nil
}

func (agg *Aggregator) makeBatch() ([]types.Task, *smt.StandardTree, error) {
	agg.taskMu.RLock()
	defer agg.taskMu.RUnlock()

	// Sort pending tasks.
	tasks := make([]types.Task, 0, len(agg.pendingTasks))
	for i, t := range agg.pendingTasks {
		tasks = append(tasks, types.Task{
			ILambadaCoprocessorTaskManagerTask: t,
			Index:                              i,
		})
	}
	taskCmp := func(t1, t2 types.Task) int {
		return cmp.Compare(t1.Index, t2.Index)
	}
	slices.SortFunc(tasks, taskCmp)

	// Build merkle tree for tasks in the batch.
	values := make([][]interface{}, len(tasks))
	leafEncodings := make([]string, len(tasks))
	for i, t := range tasks {
		taskHash, err := core.GetTaskDigest(&t.ILambadaCoprocessorTaskManagerTask)
		if err != nil {
			return nil, nil, err
		}
		values[i] = []interface{}{
			smt.SolBytes(string(taskHash[:])),
		}
		leafEncodings[i] = smt.SOL_BYTES32
	}
	merkle, err := smt.Of(values, leafEncodings)
	if err != nil {
		return nil, nil, err
	}

	return tasks, merkle, nil
}

func (agg *Aggregator) confirmBatch(
	onchain tm.ILambadaCoprocessorTaskManagerTaskBatch,
	tasks []types.Task,
	merkle *smt.StandardTree,
) types.TaskBatch {
	agg.taskMu.Lock()
	defer agg.taskMu.Unlock()

	// Update batch lookup.
	agg.batchIndex = onchain.Index
	batch := types.TaskBatch{
		ILambadaCoprocessorTaskManagerTaskBatch: onchain,
		Tasks:                                   tasks,
		Merkle:                                  merkle,
	}
	agg.batches[batch.Index] = batch

	// Update task lookup.
	for _, t := range batch.Tasks {
		delete(agg.pendingTasks, t.Index)
		agg.submittedTasks[t.Index] = t.ILambadaCoprocessorTaskManagerTask
	}

	return batch
}

func (agg *Aggregator) processTaskResponse(
	taskIndex sdktypes.TaskIndex,
	operatorID bls.OperatorId,
	response tm.ILambadaCoprocessorTaskManagerTaskResponse,
	sig bls.Signature,
) error {
	agg.taskMu.Lock()
	defer agg.taskMu.Unlock()

	// Create new lookup of task's responses if neccessary.
	if _, ok := agg.taskResponses[taskIndex]; !ok {
		agg.taskResponses[taskIndex] = make(map[sdktypes.TaskResponseDigest]tm.ILambadaCoprocessorTaskManagerTaskResponse)
	}

	// Memorize response from operator.
	responseDigest, err := core.GetTaskResponseDigest(&response)
	if err != nil {
		agg.log.Error("Failed to get task response digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}
	if _, ok := agg.taskResponses[taskIndex][responseDigest]; !ok {
		agg.taskResponses[taskIndex][responseDigest] = response
	}

	return agg.blsAggregationService.ProcessNewSignature(
		context.Background(), taskIndex, responseDigest, &sig, operatorID,
	)
}

func (agg *Aggregator) sendAggregatedResponseToContract(
	blsAggServiceResp blsagg.BlsAggregationServiceResponse,
) error {
	// TODO: check if blsAggServiceResp contains an err
	if blsAggServiceResp.Err != nil {
		agg.log.Error("BlsAggregationServiceResponse contains an error", "err", blsAggServiceResp.Err)
		// panicing to help with debugging (fail fast), but we shouldn't panic if we run this in production
		panic(blsAggServiceResp.Err)
	}
	nonSignerPubkeys := []tm.BN254G1Point{}
	for _, nonSignerPubkey := range blsAggServiceResp.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, core.ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []tm.BN254G1Point{}
	for _, quorumApk := range blsAggServiceResp.QuorumApksG1 {
		quorumApks = append(quorumApks, core.ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := tm.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(blsAggServiceResp.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(blsAggServiceResp.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: blsAggServiceResp.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             blsAggServiceResp.QuorumApkIndices,
		TotalStakeIndices:            blsAggServiceResp.TotalStakeIndices,
		NonSignerStakeIndices:        blsAggServiceResp.NonSignerStakeIndices,
	}

	agg.log.Info("Threshold reached. Sending aggregated response onchain.",
		"taskIndex", blsAggServiceResp.TaskIndex,
	)

	// Get task, batch and task response from local "storage".
	agg.taskMu.Lock()
	batch, ok := agg.batches[agg.batchIndex]
	if !ok {
		panic("batch with sepcified index does not exist")
	}
	task, taskBatchIdx, ok := batch.TaskByIndex(blsAggServiceResp.TaskIndex)
	if !ok {
		panic("batch does not contain task with specified index")
	}
	taskResponse, ok := agg.taskResponses[blsAggServiceResp.TaskIndex][blsAggServiceResp.TaskResponseDigest]
	if !ok {
		panic("response with specified digest and task index deos not exist")
	}
	agg.taskMu.Unlock()

	// Generate proof for task.
	taskProof, err := batch.Merkle.GetProofWithIndex(taskBatchIdx)
	if err != nil {
		return err
	}
	taskProofOnchain := make([][32]byte, len(taskProof))
	for i, p := range taskProof {
		taskProofOnchain[i] = [32]byte(p)
	}

	_, err = agg.avsWriter.RespondTask(
		context.Background(),
		batch.ILambadaCoprocessorTaskManagerTaskBatch,
		task.ILambadaCoprocessorTaskManagerTask,
		taskProofOnchain,
		taskResponse,
		nonSignerStakesAndSignature,
	)

	return err
}

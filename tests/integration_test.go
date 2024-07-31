package integration_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/zippiehq/cartesi-lambada-coprocessor/aggregator"
	aggtypes "github.com/zippiehq/cartesi-lambada-coprocessor/aggregator/types"
	tm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/chainio"
)

const (
	DEVNET  = "devnet"
	HOLESKY = "holesky"
	MAINNET = "mainnet"
)

var ZERO_TASK_OUTPUT_HASH [32]byte

var network = flag.String("network", DEVNET, "target network")

type task struct {
	ProgramID string `json:"programId"`
	Input     string `json:"input"`
}

type taskBatch struct {
	index          aggtypes.TaskBatchIndex
	tasks          []task
	firstTaskIndex sdktypes.TaskIndex
	lastTaskIndex  sdktypes.TaskIndex
}

func TestIntegration(t *testing.T) {
	var deploymentPath string
	switch *network {
	case DEVNET:
		deploymentPath = "../contracts/script/output/lambada_coprocessor_deployment_output_devnet.json"
	case HOLESKY:
		deploymentPath = "../contracts/script/output/lambada_coprocessor_deployment_output_holesky.json"
	case MAINNET:
		deploymentPath = "../contracts/script/output/lambada_coprocessor_deployment_output_mainnet.json"
	}

	var deployment chainio.AVSDeployment
	if err := sdkutils.ReadJsonConfig(deploymentPath, &deployment); err != nil {
		t.Fatalf("failed to read deployment output file - %s", err)
	}

	log, err := sdklogging.NewZapLogger(sdklogging.Development)
	if err != nil {
		t.Fatalf("failed to create Zap logger - %s", err)
	}

	rpcClient, err := eth.NewClient("http://127.0.0.1:8545")
	if err != nil {
		t.Fatalf("failed to create ETH RPC client - %s", err)
	}
	avsReader, err := chainio.NewAvsReader(deployment, rpcClient, log)
	if err != nil {
		t.Fatalf("failed to create avs reader - %s", err)
	}

	wsClient, err := eth.NewClient("ws://127.0.0.1:8545")
	if err != nil {
		t.Fatalf("failed to create ETH RPC client - %s", err)
	}
	avsSubscriber, err := chainio.NewAvsSubscriber(deployment, wsClient, log)
	if err != nil {
		t.Fatalf("failed to create avs subscriber - %s", err)
	}

	batches := []taskBatch{
		// batch 0
		{
			index: 0,
			tasks: []task{
				{
					ProgramID: "bafybeicdhhtwmgpnt7jugvlv3xtp2u4w4mkunpmg6txkkkjhpvnt2buyqa",
					Input:     "echo input1",
				},
				{
					ProgramID: "bafybeicdhhtwmgpnt7jugvlv3xtp2u4w4mkunpmg6txkkkjhpvnt2buyqa",
					Input:     "echo input2",
				},
			},
			firstTaskIndex: 0,
			lastTaskIndex:  1,
		},
		// batch 1
		{
			index: 1,
			tasks: []task{
				{
					ProgramID: "bafybeicdhhtwmgpnt7jugvlv3xtp2u4w4mkunpmg6txkkkjhpvnt2buyqa",
					Input:     "echo input3",
				},
			},
			firstTaskIndex: 2,
			lastTaskIndex:  2,
		},
		// batch 2
		{
			index: 2,
			tasks: []task{
				{
					ProgramID: "bafybeicdhhtwmgpnt7jugvlv3xtp2u4w4mkunpmg6txkkkjhpvnt2buyqa",
					Input:     "echo input4",
				},
				{
					ProgramID: "bafybeicdhhtwmgpnt7jugvlv3xtp2u4w4mkunpmg6txkkkjhpvnt2buyqa",
					Input:     "echo input5",
				},
				{
					ProgramID: "bafybeicdhhtwmgpnt7jugvlv3xtp2u4w4mkunpmg6txkkkjhpvnt2buyqa",
					Input:     "echo input6",
				},
			},
			firstTaskIndex: 3,
			lastTaskIndex:  5,
		},
	}

	for _, b := range batches {
		checkTaskBatch(t, b, avsReader, avsSubscriber)
	}
}

func checkTaskBatch(
	t *testing.T,
	batch taskBatch,
	avsReader *chainio.AvsReader, avsSubscriber *chainio.AvsSubscriber,
) {
	batchCh := make(chan *tm.ContractLambadaCoprocessorTaskManagerTaskBatchRegistered)
	if _, err := avsSubscriber.SubscribeToNewBatches(batchCh); err != nil {
		t.Fatalf("failed to subscribe to new task batches - %s", err)
	}

	respCh := make(chan *tm.ContractLambadaCoprocessorTaskManagerTaskResponded)
	if _, err := avsSubscriber.SubscribeToTaskResponses(respCh); err != nil {
		t.Fatalf("failed to subscribe to task responses - %s", err)
	}

	aggTaskIdx, err := submitTasks(batch.tasks)
	if err != nil {
		t.Fatalf("failed to submit tasks - %s", err)
	}
	assert.Equal(t, batch.lastTaskIndex, aggTaskIdx, "invalid aggregator task index for batch %d", batch.index)

	// Validate onchain batch.
	select {
	case onchainBatch := <-batchCh:
		// Validate batch index.
		nextBatchIdx, err := avsReader.Bindings.TaskManager.NextBatchIndex(&bind.CallOpts{})
		if err != nil {
			t.Fatalf("failed to fetch next batch index - %s", err)
		}
		assert.Equal(t, batch.index, onchainBatch.Batch.Index, "invalid onchain batch index for batch %d", batch.index)
		assert.Equal(t, batch.index+1, nextBatchIdx, "invalid onchain next batch index for batch %d", batch.index)

		// Validate merkle root.
		batchMerkleRoot, err := batchMerkleRoot(batch)
		if err != nil {
			t.Fatalf("failed to compute merkle root for task batch - %s", err)
		}
		assert.Equal(t, batchMerkleRoot, onchainBatch.Batch.MerkeRoot, "invalid onchain batch merkle root for batch %d", batch.index)

		// TODO: validate block, quorumNumbers and quroumThresholdPrecentage?

		// Validate batch mapping.
		batchHash, err := core.GetTaskBatchDigest(&onchainBatch.Batch)
		if err != nil {
			t.Fatalf("failed to compute task batch hash - %s", err)
		}
		onchainBatchHash, err := avsReader.Bindings.TaskManager.AllBatchHashes(&bind.CallOpts{}, batch.index)
		if err != nil {
			t.Fatalf("failed to fetch batch hash - %s", err)
		}
		assert.Equal(t, batchHash, onchainBatchHash, "invalid onchain batch hash for batch %d", batch.index)

	case <-time.After(15 * time.Second):
		t.Fatalf("failed to get new batch in 15 seconds")
	}

	// Validate all onchain task responses.
	for range batch.tasks {
		select {
		case <-respCh:
		case <-time.After(300 * time.Second):
			t.Fatalf("failed to get task response in 300 seconds")
		}
	}
	// Validate task response mapping.
	for _, task := range batch.tasks {
		inputHash := core.Keccack256([]byte(task.Input))
		outputHash, err := avsReader.Bindings.TaskManager.GetTaskResponseHash(&bind.CallOpts{}, batch.index, []byte(task.ProgramID), inputHash)
		if err != nil {
			t.Fatalf("failed to fetch task output hash - %s", err)
		}
		assert.NotEqual(
			t,
			ZERO_TASK_OUTPUT_HASH,
			outputHash,
			"can not find task output hash for batch %d and program %s",
			batch.index,
			task.ProgramID,
		)
	}
}

func submitTasks(tasks []task) (sdktypes.TaskIndex, error) {
	body, err := json.Marshal(tasks)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post("http://localhost:8080/createTask", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	respJSON := struct {
		TaskIndex sdktypes.TaskIndex `json:"taskIndex"`
	}{}
	if err = json.Unmarshal(respData, &respJSON); err != nil {
		return 0, err
	}

	return respJSON.TaskIndex, nil
}

func batchMerkleRoot(batch taskBatch) ([32]byte, error) {
	// Tasks in aggregator must be in order
	aggTasks := make([]aggtypes.Task, len(batch.tasks))
	for i, t := range batch.tasks {
		aggTasks[i] = aggtypes.Task{
			ILambadaCoprocessorTaskManagerTask: tm.ILambadaCoprocessorTaskManagerTask{
				ProgramId: []byte(t.ProgramID),
				InputHash: core.Keccack256([]byte(t.Input)),
			},
			Input: []byte(t.Input),
			Index: batch.firstTaskIndex + uint32(i),
		}
	}

	_, stm, err := aggregator.BuildTaskBatchMerkle(aggTasks)
	if err != nil {
		var nullRoot [32]byte
		return nullRoot, err
	}

	return [32]byte(stm.GetRoot()), nil
}

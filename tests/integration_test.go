package integration_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/zippiehq/cartesi-lambada-coprocessor/aggregator"
	aggtypes "github.com/zippiehq/cartesi-lambada-coprocessor/aggregator/types"
	tm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/chainio"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/config"
)

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
	config, err := config.NewConfig(
		"./nodes/aggregator/aggregator.yaml",
		"../contracts/script/output/31337/lambada_coprocessor_avs_deployment_output.json",
		"0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
	)
	if err != nil {
		t.Fatalf("failed to create avs config - %s", err)
	}

	avsReader, err := chainio.BuildAvsReaderFromConfig(config)
	if err != nil {
		t.Fatalf("failed to create avs reader - %s", err)
	}
	avsSubscriber, err := chainio.BuildAvsSubscriberFromConfig(config)
	if err != nil {
		t.Fatalf("failed to create avs subscriber - %s", err)
	}

	batches := []taskBatch{
		// batch 0
		{
			index: 0,
			tasks: []task{
				{
					ProgramID: "program1",
					Input:     "input1",
				},
				{
					ProgramID: "program1",
					Input:     "input2",
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
					ProgramID: "program2",
					Input:     "input1",
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
					ProgramID: "program1",
					Input:     "input3",
				},
				{
					ProgramID: "program2",
					Input:     "input4",
				},
				{
					ProgramID: "program3",
					Input:     "input5",
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
		nextBatchIdx, err := avsReader.AvsManagersBindings.TaskManager.NextBatchIndex(&bind.CallOpts{})
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
		onchainBatchHash, err := avsReader.AvsManagersBindings.TaskManager.AllBatchHashes(&bind.CallOpts{}, batch.index)
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
		onchainHash, err := core.GetTaskResponseMetadataDigest(batch.index, []byte(task.ProgramID), []byte(task.Input))
		if err != nil {
			t.Fatalf("failed to compute task response metadata hash -%s", err)
		}
		respExists, err := avsReader.AvsManagersBindings.TaskManager.AllTaskResponses(&bind.CallOpts{}, onchainHash)
		if err != nil {
			t.Fatalf("failed to fetch task response flag - %s", err)
		}
		assert.True(t, respExists, "invalid onchain task response flag for batch %d", batch.index)
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
				Input:     []byte(t.Input),
			},
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

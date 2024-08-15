package core

import (
	"cmp"
	"encoding/hex"
	"fmt"
	"slices"

	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

type TaskBatchIndex = uint32

type TaskDigest [32]byte

type Task struct {
	Index     sdktypes.TaskIndex `json:"index"`
	ProgramID []byte             `json:"programID"`
	Input     []byte             `json:"input"`
	InputHash []byte             `json:"inputHash"`
}

type TaskBatch struct {
	Index                     uint32            `json:"index"`
	BlockNumber               uint32            `json:"blockNumber"`
	QuorumNumbers             []byte            `json:"quroumNumbers"`
	QuorumThresholdPercentage uint32            `json:"quroumThresholdPercentage"`
	Tasks                     []Task            `json:"tasks"`
	Merkle                    *smt.StandardTree `json:"merkle"`
}

type TaskResponse struct {
	TaskIndex  uint32              `json:"taskIndex"`
	OperatorID sdktypes.OperatorId `json:"operatorID"`
	OutputHash []byte              `json:"outputHash"`
	ResultCID  []byte              `json:"resultCID"`
}

func BuildTaskBatchMerkle(tasks []Task) ([]Task, *smt.StandardTree, error) {
	taskCmp := func(t1, t2 Task) int {
		return cmp.Compare(t1.Index, t2.Index)
	}
	slices.SortFunc(tasks, taskCmp)

	// Build merkle tree for tasks in the batch.
	values := make([][]interface{}, len(tasks))
	for i, t := range tasks {
		values[i] = []interface{}{
			smt.SolBytes(hex.EncodeToString(t.ProgramID)),
			smt.SolBytes(hex.EncodeToString(t.InputHash)),
		}
	}

	leafEncodings := []string{
		smt.SOL_BYTES,
		smt.SOL_BYTES,
	}
	merkle, err := smt.Of(values, leafEncodings)
	if err != nil {
		return nil, nil, err
	}

	return tasks, merkle, nil
}

func TaskProof(batch TaskBatch, taskIdx sdktypes.TaskIndex) ([][]byte, error) {
	taskCmp := func(t1, t2 Task) int {
		return cmp.Compare(t1.Index, t2.Index)
	}
	slices.SortFunc(batch.Tasks, taskCmp)

	for i, t := range batch.Tasks {
		if t.Index == taskIdx {
			return batch.Merkle.GetProofWithIndex(i)
		}
	}

	return nil, fmt.Errorf("batch does not contain task with index %d", taskIdx)
}

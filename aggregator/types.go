package aggregator

import (
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
	Index                     uint32 `json:"index"`
	BlockNumber               uint32 `json:"blockNumber"`
	QuorumNumbers             []byte `json:"quroumNumbers"`
	QuorumThresholdPercentage uint32 `json:"quorumThresholdPrecentage"`
	SubmittedTasks            []Task `json:"submittedTasks"`
	Merkle                    []byte `json:"merkle"`
}

type TaskResponse struct {
	ResultCID  []byte   `json:"resultCID"`
	OutputHash [32]byte `json:"outputHash"`
}

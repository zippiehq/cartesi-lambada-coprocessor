package aggregator

import (
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

package types

import (
	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/ethereum/go-ethereum/common"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	tm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
)

// all operators in quorum0 must sign the task response in order for it to be accepted
// TODO: our contracts require uint8 but right now sdktypes.QuorumThresholdPercentage is uint8 prob need to update our inc-sq contracts to use uint8 as well?
const QUORUM_THRESHOLD_NUMERATOR = sdktypes.QuorumThresholdPercentage(100)
const QUORUM_THRESHOLD_DENOMINATOR = sdktypes.QuorumThresholdPercentage(100)

const QUERY_FILTER_FROM_BLOCK = uint64(1)

// we only use a single quorum (quorum 0) for incredible squaring
var QUORUM_NUMBERS = sdktypes.QuorumNums{0}

type BlockNumber = uint32

type OperatorInfo struct {
	OperatorPubkeys sdktypes.OperatorPubkeys
	OperatorAddr    common.Address
}

type TaskBatchIndex = uint32

type TaskDigest [32]byte

type Task struct {
	tm.ILambadaCoprocessorTaskManagerTask
	Index sdktypes.TaskIndex
}

type TaskBatch struct {
	tm.ILambadaCoprocessorTaskManagerTaskBatch
	Tasks  []Task
	Merkle *smt.StandardTree
}

func (b *TaskBatch) TaskByIndex(idx sdktypes.TaskIndex) (Task, int, bool) {
	for i, t := range b.Tasks {
		if t.Index == idx {
			return t, i, true
		}
	}
	return Task{}, 0, false
}

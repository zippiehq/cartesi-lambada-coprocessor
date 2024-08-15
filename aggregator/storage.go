package aggregator

import (
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/zippiehq/cartesi-lambada-coprocessor/core"
)

type Storage interface {
	AddPendingTask(core.Task) (sdktypes.TaskIndex, error)
	AllPendingTasks() ([]core.Task, error)

	AddTaskBatch(core.TaskBatch) error
	TaskBatch(core.TaskBatchIndex) (core.TaskBatch, error)

	SubmittedTask(sdktypes.TaskIndex) (core.Task, core.TaskBatchIndex, error)

	AddTaskResponse(core.TaskResponse) error
	TaskResponses(sdktypes.TaskIndex) ([]core.TaskResponse, error)
}

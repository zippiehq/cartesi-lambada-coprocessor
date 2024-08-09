package aggregator

import sdktypes "github.com/Layr-Labs/eigensdk-go/types"

type Storage interface {
	AddPendingTask(Task) error
	PendingTask(sdktypes.TaskIndex) (Task, error)
	AllPendingTasks() ([]Task, error)

	AddTaskBatch(TaskBatch) error
	TaskBatch(TaskBatchIndex) (TaskBatch, error)
	AllTaskBatches() ([]TaskBatch, error)

	AddTaskResponse(sdktypes.TaskIndex, TaskResponse) error
	TaskResponses(sdktypes.TaskIndex) (map[sdktypes.TaskResponseDigest]TaskResponse, error)
}

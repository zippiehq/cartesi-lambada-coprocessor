package aggregator

import sdktypes "github.com/Layr-Labs/eigensdk-go/types"

type Storage interface {
	AddPendingTask(Task) (sdktypes.TaskIndex, error)
	AllPendingTasks() ([]Task, error)

	AddTaskBatch(TaskBatch) error
	TaskBatch(TaskBatchIndex) (TaskBatch, error)

	SubmittedTask(sdktypes.TaskIndex) (Task, TaskBatchIndex, error)

	AddTaskResponse(TaskResponse) error
	TaskResponses(sdktypes.TaskIndex) ([]TaskResponse, error)
}

package aggregator

import sdktypes "github.com/Layr-Labs/eigensdk-go/types"

type MySqlStoageConfig struct {
}

type MySqlStorage struct {
}

func NewMySqlStorage(cfg MySqlStoageConfig) (*MySqlStorage, error) {
	return nil, nil
}

func (s *MySqlStorage) AddPedningTask(i sdktypes.TaskIndex) error {
	return nil
}

func (s *MySqlStorage) PendingTask(sdktypes.TaskIndex) (Task, error) {
	return Task{}, nil
}

func (s *MySqlStorage) AllPendingTasks() ([]Task, error) {
	return nil, nil
}

func (s *MySqlStorage) AddTaskBatch(TaskBatch) error {
	return nil
}

func (s *MySqlStorage) TaskBatch(TaskBatchIndex) (TaskBatch, error) {
	return TaskBatch{}, nil
}

func (s *MySqlStorage) AllTaskBatches() ([]TaskBatch, error) {
	return nil, nil
}

func (s *MySqlStorage) AddTaskResponse(sdktypes.TaskIndex, TaskResponse) error {
	return nil
}

func (s *MySqlStorage) TaskResponses(sdktypes.TaskIndex) (map[sdktypes.TaskResponseDigest]TaskResponse, error) {
	return nil, nil
}

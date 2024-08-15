package aggregator

import (
	"database/sql"
	"encoding/base64"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

type MySqlStorageConfig struct {
	Address  string
	Database string
	User     string
	Password string
}

type MySqlStorage struct {
	db *sql.DB
}

func NewMySqlStorage(cfg MySqlStorageConfig) (*MySqlStorage, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.User, cfg.Password, cfg.Address, cfg.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	s := &MySqlStorage{
		db: db,
	}

	return s, nil
}

func (s *MySqlStorage) AddPendingTask(t core.Task) (sdktypes.TaskIndex, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}

	row := tx.QueryRow("SELECT * FROM tasks ORDER BY task_index DESC LIMIT 1")
	lastTask, _, err := readTask(row)
	var taskIndex sdktypes.TaskIndex
	if err != nil {
		if err == sql.ErrNoRows {
			taskIndex = 1
		} else {
			return 0, err
		}
	} else {
		taskIndex = lastTask.Index + 1
	}

	stmt, err := tx.Prepare("INSERT INTO tasks VALUES(NULL, ?, ?, ?, NULL)")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if _, err = stmt.Exec(
		base64.StdEncoding.EncodeToString(t.ProgramID),
		base64.StdEncoding.EncodeToString(t.Input),
		base64.StdEncoding.EncodeToString(t.InputHash),
	); err != nil {
		tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return taskIndex, nil
}

func (s *MySqlStorage) AllPendingTasks() ([]core.Task, error) {
	rows, err := s.db.Query("SELECT * FROM tasks WHERE batch_index IS NULL ORDER BY task_index ASC")
	if err != nil {
		return nil, err
	}

	tasks := make([]core.Task, 0)
	for rows.Next() {
		t, _, err := readTask(rows)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (s *MySqlStorage) AddTaskBatch(b core.TaskBatch) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Write new task batch.
	stmt, err := tx.Prepare("INSERT INTO task_batches VALUES(?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err = stmt.Exec(
		b.Index,
		b.BlockNumber,
		base64.StdEncoding.EncodeToString(b.QuorumNumbers),
		b.QuorumThresholdPercentage,
	); err != nil {
		tx.Rollback()
		return err
	}

	// Mark existing tasks as submitted.
	for _, t := range b.Tasks {
		stmt, err := tx.Prepare("UPDATE tasks SET batch_index=? WHERE task_index=?")
		if err != nil {
			tx.Rollback()
			return err
		}
		if _, err = stmt.Exec(b.Index, t.Index); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *MySqlStorage) TaskBatch(i core.TaskBatchIndex) (core.TaskBatch, error) {
	br := s.db.QueryRow("SELECT * FROM task_batches WHERE batch_index=?", i)
	b, err := readTaskBatch(br)
	if err != nil {
		return core.TaskBatch{}, err
	}

	tr, err := s.db.Query("SELECT * from tasks WHERE batch_index=? ORDER BY task_index ASC", i)
	if err != nil {
		return core.TaskBatch{}, err
	}

	for tr.Next() {
		t, _, err := readTask(tr)
		if err != nil {
			return core.TaskBatch{}, err
		}
		b.Tasks = append(b.Tasks, t)
	}

	if _, b.Merkle, err = core.BuildTaskBatchMerkle(b.Tasks); err != nil {
		return core.TaskBatch{}, err
	}

	return b, nil
}

func (s *MySqlStorage) SubmittedTask(i sdktypes.TaskIndex) (core.Task, core.TaskBatchIndex, error) {
	r := s.db.QueryRow("SELECT * FROM tasks WHERE task_index=? AND batch_index IS NOT NULL", i)
	t, batchIdx, err := readTask(r)
	if err != nil {
		return core.Task{}, 0, err
	}

	return t, batchIdx, nil
}

func (s *MySqlStorage) AddTaskResponse(r core.TaskResponse) error {
	stmt, err := s.db.Prepare("INSERT INTO task_responses VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(
		r.TaskIndex,
		base64.StdEncoding.EncodeToString(r.OperatorID[:]),
		base64.StdEncoding.EncodeToString(r.OutputHash),
		base64.StdEncoding.EncodeToString(r.ResultCID),
	); err != nil {
		return err
	}

	return nil
}

func (s *MySqlStorage) TaskResponses(ti sdktypes.TaskIndex) ([]core.TaskResponse, error) {
	rows, err := s.db.Query("SELECT * FROM task_responses WHERE task_index=? ORDER BY task_index ASC", ti)
	if err != nil {
		return nil, err
	}

	responses := make([]core.TaskResponse, 0)
	for rows.Next() {
		r, err := readTaskResponse(rows)
		if err != nil {
			return nil, err
		}
		responses = append(responses, r)
	}

	return responses, nil
}

type rowScanner interface {
	Scan(dest ...any) error
}

func readTask(rs rowScanner) (core.Task, core.TaskBatchIndex, error) {
	var (
		t            core.Task
		programIDStr string
		inputStr     string
		inputHashStr string
		batchIdx     sql.NullInt32
	)
	err := rs.Scan(&t.Index, &programIDStr, &inputStr, &inputHashStr, &batchIdx)
	if err != nil {
		return core.Task{}, 0, err
	}

	if t.ProgramID, err = base64.StdEncoding.DecodeString(programIDStr); err != nil {
		return core.Task{}, 0, err
	}
	if t.Input, err = base64.StdEncoding.DecodeString(inputStr); err != nil {
		return core.Task{}, 0, err
	}
	if t.InputHash, err = base64.StdEncoding.DecodeString(inputHashStr); err != nil {
		return core.Task{}, 0, err
	}

	return t, uint32(batchIdx.Int32), nil
}

func readTaskBatch(rs rowScanner) (core.TaskBatch, error) {
	var (
		b                core.TaskBatch
		quorumNumbersStr string
	)
	err := rs.Scan(
		&b.Index,
		&b.BlockNumber,
		&quorumNumbersStr,
		&b.QuorumThresholdPercentage)
	if err != nil {
		return core.TaskBatch{}, err
	}

	if b.QuorumNumbers, err = base64.StdEncoding.DecodeString(quorumNumbersStr); err != nil {
		return core.TaskBatch{}, err
	}
	b.Tasks = make([]core.Task, 0)

	return b, nil
}

func readTaskResponse(rs rowScanner) (core.TaskResponse, error) {
	var (
		r             core.TaskResponse
		operatorIDStr string
		resultCIDStr  string
		outputHashStr string
	)

	err := rs.Scan(&r.TaskIndex, &operatorIDStr, &outputHashStr, &resultCIDStr)
	if err != nil {
		return core.TaskResponse{}, err
	}

	operatorID, err := base64.StdEncoding.DecodeString(operatorIDStr)
	if err != nil {
		return core.TaskResponse{}, err
	}
	r.OperatorID = sdktypes.Bytes32(operatorID)

	if r.OutputHash, err = base64.StdEncoding.DecodeString(outputHashStr); err != nil {
		return core.TaskResponse{}, err
	}
	if r.ResultCID, err = base64.StdEncoding.DecodeString(resultCIDStr); err != nil {
		return core.TaskResponse{}, err
	}

	return r, nil
}

package aggregator

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/zippiehq/cartesi-lambada-coprocessor/aggregator/types"
)

func (agg *Aggregator) startAPIServer() error {
	http.HandleFunc("/createTask", agg.createTask)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return nil
}

func (agg *Aggregator) createTask(w http.ResponseWriter, r *http.Request) {
	taskData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tasks := make(
		[]struct {
			ProgramID string `json:"programId"`
			Input     string `json:"input"`
		},
		0,
	)
	if err = json.Unmarshal(taskData, &tasks); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	batchIndex := agg.getBatchIndex()
	var taskIndex sdktypes.TaskIndex
	for _, t := range tasks {
		if taskIndex, err = agg.addTask([]byte(t.ProgramID), []byte(t.Input)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	resp := struct {
		BatchIndex types.TaskBatchIndex `json:"batchIndex"`
		TaskIndex  sdktypes.TaskIndex   `json:"taskIndex"`
	}{
		BatchIndex: batchIndex,
		TaskIndex:  taskIndex,
	}
	respData, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respData)
}

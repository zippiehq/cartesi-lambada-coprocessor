package aggregator

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

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

	task := struct {
		ProgramID string `json:"programId"`
		Input     string `json:"input"`
	}{}
	if err = json.Unmarshal(taskData, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskIndex, err := agg.addTask([]byte(task.ProgramID), []byte(task.Input))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		TaskIndex types.TaskBatchIndex `json:"taskIndex"`
	}{
		TaskIndex: taskIndex,
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

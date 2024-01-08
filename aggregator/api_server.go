package aggregator

import (
	"encoding/json"
	"io"
	"log"
	"math/big"
	"net/http"

	"github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
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
		Number int64 `json:"number"`
	}{}
	if err = json.Unmarshal(taskData, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskIndex, err := agg.sendNewTask(big.NewInt(task.Number))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		TaskIndex types.TaskIndex `json:"taskIndex"`
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

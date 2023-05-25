package task

import (
	"encoding/json"
	taskUsecase "go-ca/internal/app/application/usecase/task"
	"net/http"
)

type TasksController struct {
	GetTasksInputPort taskUsecase.GetTasksInputPort
}

type Response struct {
	TaskId       uint64 `json:"task_id"`
	TaskName     string `json:"task_name"`
	Reward       uint64 `json:"reward"`
	ContractName string `json:"contract_name"`
}

func NewTasksController(GetTasksInputPort taskUsecase.GetTasksInputPort) TasksController {
	return TasksController{GetTasksInputPort: GetTasksInputPort}
}

func (c *TasksController) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := c.GetTasksInputPort.GetTasks()

	// 構造体をレスポンスに変換
	response := make([]*Response, len(tasks))
	for i, t := range tasks {
		response[i] = &Response{
			TaskId:       t.TaskId,
			TaskName:     t.TaskName,
			Reward:       t.Reward,
			ContractName: t.ContractName,
		}
	}

	// 構造体をJSONに変換
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスを設定
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

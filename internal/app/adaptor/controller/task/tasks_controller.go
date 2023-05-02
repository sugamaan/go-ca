package task

import (
	"encoding/json"
	"fmt"
	"go-ca/internal/app/application/usecase/task"
	"net/http"
)

type TasksController struct {
	GetTasksInputPort task.GetTasksInputPort
}

type Response struct {
	TaskId uint64 `json:"task_id"`
	Name   string `json:"name"`
	Reward uint64 `json:"reward"`
}

func NewTasksController(GetTasksInputPort task.GetTasksInputPort) TasksController {
	return TasksController{GetTasksInputPort: GetTasksInputPort}
}

func (c *TasksController) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := c.GetTasksInputPort.GetTasks()
	fmt.Println("c.GetTasksInputPort.GetTasks()")
	fmt.Printf("%#v\n", tasks)

	// 構造体をレスポンスに変換
	response := make([]*Response, len(tasks))
	for i, t := range tasks {
		response[i] = &Response{
			TaskId: t.GetTaskId(),
			Name:   t.GetName(),
			Reward: uint64(t.GetReward()),
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

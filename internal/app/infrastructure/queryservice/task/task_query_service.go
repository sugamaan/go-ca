package task

import "go-ca/internal/app/application/usecase/task"

type TaskQueryService struct {
}

func NewTaskQueryService() *TaskQueryService {
	return &TaskQueryService{}
}

func (s TaskQueryService) GetTaskByUserId() (task.GetTaskByUserIdDto, error) {
	task := task.GetTaskByUserIdDto{
		TaskId:   1,
		TaskName: "テスト",
	}
	return task, nil
}

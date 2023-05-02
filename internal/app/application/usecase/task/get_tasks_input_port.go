package task

import "go-ca/internal/app/domain/task/entity"

type GetTasksInputPort interface {
	GetTasks() []*entity.Task
}

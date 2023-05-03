package task

import (
	"go-ca/internal/app/domain/task/entity"
	"go-ca/internal/app/domain/task/repository"
)

type GetTasksUsecase struct {
	tasksRepository repository.TasksRepository
}

type GetTasksQueryService interface {
	GetTasks() []GetTasksDto
}

type GetTasksDto struct {
	TaskId uint64
	Name   string
	Reward uint64
}

func NewGetTasksUsecase(tasksRepository repository.TasksRepository) GetTasksUsecase {
	return GetTasksUsecase{tasksRepository: tasksRepository}
}

func (u *GetTasksUsecase) GetTasks() []*entity.Task {
	// DBからタスク一覧を取得
	tasks, err := u.tasksRepository.GetAllTasks()
	if err != nil {
		// TODO エラー処理
		panic(err)
	}

	// UI層へ受け渡す
	return tasks
}

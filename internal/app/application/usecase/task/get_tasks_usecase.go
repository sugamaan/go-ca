package task

import (
	taskDomain "go-ca/internal/app/domain/task"
)

type GetTasksUsecase struct {
	tasksRepository taskDomain.TasksRepository
}

func NewGetTasksUsecase(tasksRepository taskDomain.TasksRepository) GetTasksUsecase {
	return GetTasksUsecase{tasksRepository: tasksRepository}
}

func (u *GetTasksUsecase) GetTasks() []*taskDomain.Task {
	// DBからタスク一覧を取得
	tasks, err := u.tasksRepository.GetAllTasks()
	if err != nil {
		// TODO エラー処理
		panic(err)
	}

	// UI層へ受け渡す
	return tasks
}

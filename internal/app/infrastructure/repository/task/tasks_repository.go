package task

import (
	"go-ca/internal/app/domain/task/repository"
)

type TasksRepository struct {
}

func NewTasksRepository() TasksRepository {
	return TasksRepository{}
}

func (r *TasksRepository) GetAllTasks() ([]repository.GetTasksDataModel, error) {
	// TODO DBアクセスを行いデータを取得
	tasksDataModel := []repository.GetTasksDataModel{
		{
			TaskId: 1,
			Name:   "テスト",
			Reward: 100,
		},
		{
			TaskId: 2,
			Name:   "テスト2",
			Reward: 200,
		},
	}

	return tasksDataModel, nil
}

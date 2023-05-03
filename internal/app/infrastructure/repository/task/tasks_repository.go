package task

import (
	"go-ca/internal/app/domain/task/entity"
	"go-ca/internal/app/domain/task/valueobject"
)

type GetTasksDataModel struct {
	TaskId uint64 `db:"task_id"`
	Name   string `db:"name"`
	Reward uint64 `db:"reward"`
}

type TasksRepository struct {
}

func NewTasksRepository() TasksRepository {
	return TasksRepository{}
}

func (r *TasksRepository) GetAllTasks() ([]*entity.Task, error) {
	// TODO DBアクセスを行いデータを取得
	tasksDataModel := []GetTasksDataModel{
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

	// データモデルからドメインへ詰め替える
	// データモデルはDBとの接続用の構造体
	tasks := make([]*entity.Task, len(tasksDataModel))
	for i, t := range tasksDataModel {
		reward, err := valueobject.NewReward(t.Reward)
		if err != nil {
			return nil, err
		}
		tasks[i], err = entity.NewTask(t.TaskId, t.Name, reward)
		if err != nil {
			return nil, err
		}
	}

	return tasks, nil
}

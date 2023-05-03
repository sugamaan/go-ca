package task

import (
	"go-ca/internal/app/domain/task/entity"
	"go-ca/internal/app/domain/task/valueobject"
)

type DataModel struct {
	TaskId uint64 `db:"task_id"`
	Name   string `db:"name"`
	Reward uint64 `db:"reward"`
}

type TasksRepository struct {
	// TODO DBアクセスに関する構造体を追加
}

func NewTasksRepository() TasksRepository {
	return TasksRepository{}
}

func (r *TasksRepository) GetAllTasks() ([]*entity.Task, error) {
	// TODO DBアクセスを行いデータを取得
	tasksDataModel := []DataModel{
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
	for i, m := range tasksDataModel {
		task, err := m.toTask()
		if err != nil {
			return nil, err
		}
		tasks[i] = task
	}

	return tasks, nil
}

func (m DataModel) toTask() (*entity.Task, error) {
	reward, err := valueobject.NewReward(m.Reward)
	if err != nil {
		return nil, err
	}
	task, err := entity.NewTask(m.TaskId, m.Name, reward)
	if err != nil {
		return nil, err
	}
	return task, nil
}

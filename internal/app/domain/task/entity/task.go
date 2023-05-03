package entity

import (
	"errors"
	"go-ca/internal/app/domain/task/valueobject"
)

type Task struct {
	taskId uint64
	name   string
	reward valueobject.Reward
}

func NewTask(taskId uint64, name string, reward valueobject.Reward) (*Task, error) {
	if name == "" || len(name) > 50 {
		return &Task{}, errors.New("タスク名が不正です")
	}

	return &Task{
		taskId: taskId,
		name:   name,
		reward: reward,
	}, nil
}

func (t Task) TaskId() uint64 {
	return t.taskId
}

func (t Task) Name() string {
	return t.name
}

func (t Task) Reward() valueobject.Reward {
	return t.reward
}

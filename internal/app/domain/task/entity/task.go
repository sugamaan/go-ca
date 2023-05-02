package entity

import (
	"errors"
	"go-ca/internal/app/domain/task/value_object"
)

type Task struct {
	taskId uint64
	name   string
	reward value_object.Reward
}

func NewTask(taskId uint64, name string, reward value_object.Reward) (*Task, error) {
	if name == "" || len(name) > 50 {
		return &Task{}, errors.New("タスク名が不正です")
	}

	return &Task{
		taskId: taskId,
		name:   name,
		reward: reward,
	}, nil
}

func (t Task) GetTaskId() uint64 {
	return t.taskId
}

func (t Task) GetName() string {
	return t.name
}

func (t Task) GetReward() value_object.Reward {
	return t.reward
}

package task

import (
	"errors"
)

type Task struct {
	taskId uint64
	name   string
	reward Reward
}

func NewTask(taskId uint64, name string, reward Reward) (*Task, error) {
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

func (t Task) Reward() Reward {
	return t.reward
}

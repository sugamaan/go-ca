package task

import (
	"errors"
	"fmt"
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

	// テスト 同パッケージ内でprivateな構造体を呼び出してみる
	// TODO 以下を呼べなくする
	// reward = RewardImpl{value: 100} syntax error
	//reward2 := Reward{} Invalid composite literal type: Reward
	// しかし、rewardImplを直接呼び出せて、rewardとして返せてしまう
	a := rewardImpl{value: 100}
	fmt.Println(a)

	return &Task{
		taskId: taskId,
		name:   name,
		reward: a,
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

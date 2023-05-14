package task

type Task struct {
	taskId uint64
	name   Name
	reward Reward
}

func NewTask(taskId uint64, name Name, reward Reward) (*Task, error) {
	return &Task{
		taskId: taskId,
		name:   name,
		reward: reward,
	}, nil
}

func (t Task) TaskId() uint64 {
	return t.taskId
}

func (t Task) Name() Name {
	return t.name
}

func (t Task) Reward() Reward {
	return t.reward
}

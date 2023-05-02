package repository

type GetTasksDataModel struct {
	TaskId uint64 `db:"task_id"`
	Name   string `db:"name"`
	Reward uint64 `db:"reward"`
}

type TasksRepository interface {
	GetAllTasks() ([]GetTasksDataModel, error)
}

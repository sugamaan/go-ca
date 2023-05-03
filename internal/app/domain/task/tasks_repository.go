package task

type TasksRepository interface {
	GetAllTasks() ([]*Task, error)
}

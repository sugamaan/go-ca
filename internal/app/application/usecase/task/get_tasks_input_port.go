package task

type GetTasksInputPort interface {
	GetTasks() []GetTasksUsecaseDto
}

package task

type GetTasksInputPort interface {
	GetTasks() ([]GetTasksContainContractDto, GetTaskByUserIdDto)
}

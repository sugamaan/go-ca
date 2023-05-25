package task

type TasksQueryService interface {
	GetTasksContainContract() ([]GetTasksUsecaseDto, error)
}

type GetTasksUsecase struct {
	//tasksRepository   taskDomain.TasksRepository
	tasksQueryService TasksQueryService
}

// メモ：複数の集約をまたぐDTO
type GetTasksUsecaseDto struct {
	TaskId       uint64 `db:"task_id"`
	TaskName     string `db:"task_name"`
	Reward       uint64 `db:"reward"`
	ContractName string `db:"contract_name"`
}

func NewGetTasksUsecase(tasksQueryService TasksQueryService) GetTasksUsecase {
	return GetTasksUsecase{tasksQueryService: tasksQueryService}
}

func (u *GetTasksUsecase) GetTasks() []GetTasksUsecaseDto {
	// DBからタスク一覧を取得
	tasks, err := u.tasksQueryService.GetTasksContainContract()
	if err != nil {
		// TODO エラー処理
		panic(err)
	}

	// UI層へ受け渡す
	return tasks
}

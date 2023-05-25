package task

type TaskQueryService interface {
	GetTasksContainContract() ([]GetTasksContainContractDto, error)
	GetTaskByUserId() (GetTaskByUserIdDto, error)
}

type GetTasksUsecase struct {
	taskQueryService TaskQueryService
}

// GetTasksContainContractDto メモ：複数の集約をまたぐDTO
type GetTasksContainContractDto struct {
	TaskId       uint64 `db:"task_id"`
	TaskName     string `db:"task_name"`
	Reward       uint64 `db:"reward"`
	ContractName string `db:"contract_name"`
}

// GetTaskByUserIdDto メモ：1つの集約ならDTOではなくドメインを使っても良さそうだな...。
type GetTaskByUserIdDto struct {
	TaskId   uint64 `db:"task_id"`
	TaskName string `db:"task_name"`
}

func NewGetTasksUsecase(taskQueryService TaskQueryService) GetTasksUsecase {
	return GetTasksUsecase{taskQueryService: taskQueryService}
}

func (u *GetTasksUsecase) GetTasks() []GetTasksContainContractDto {
	// DBからタスク一覧を取得
	tasks, err := u.taskQueryService.GetTasksContainContract()
	if err != nil {
		// TODO エラー処理
		panic(err)
	}

	// DBから単一のタスクを取得
	_, _ = u.taskQueryService.GetTaskByUserId()

	// UI層へ受け渡す
	return tasks
}

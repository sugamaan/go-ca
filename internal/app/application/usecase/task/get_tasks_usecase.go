package task

import (
	contractDomain "go-ca/internal/app/domain/contract"
	taskDomain "go-ca/internal/app/domain/task"
)

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

func (u *GetTasksUsecase) GetTasks() ([]GetTasksContainContractDto, *taskDomain.Task) {
	// DBからタスク一覧を取得
	tasks, err := u.taskQueryService.GetTasksContainContract()
	if err != nil {
		// TODO エラー処理
		panic(err)
	}

	// DBから単一のタスクを取得
	task, err := u.taskQueryService.GetTaskByUserId()
	if err != nil {
		// TODO エラー処理
		panic(err)
	}

	// まずは最適化される前のコードを記述する
	// DTOからdomainへ詰め替える
	// DBから契約タイプを取得したと仮定する
	contractType := uint32(1)
	contract, err := contractDomain.NewContract(contractType)
	if err != nil {
		// TODO エラー処理
		panic(err)
	}
	reward, err := taskDomain.NewReward(taskDomain.DefaultReward, contract)
	if err != nil {
		// TODO エラー処理
		panic(err)
	}
	name, err := taskDomain.NewName(task.TaskName, contract)
	if err != nil {
		// TODO エラー処理
		panic(err)
	}
	taskDomain, err := taskDomain.NewTask(task.TaskId, name, reward)
	if err != nil {
		// TODO エラー処理
		panic(err)
	}

	// UI層へ受け渡す
	return tasks, taskDomain
}

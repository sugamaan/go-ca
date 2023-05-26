package task

import "go-ca/internal/app/application/usecase/task"

func (s TaskQueryService) GetTasksContainContract() ([]task.GetTasksContainContractDto, error) {
	// TODO DBアクセスを行いデータを取得
	// 複数の集約を跨ぐ
	getTasksUsecaseDto := []task.GetTasksContainContractDto{
		{
			TaskId:       1,
			TaskName:     "テスト",
			Reward:       100,
			ContractName: "無料",
		},
		{
			TaskId:       2,
			TaskName:     "テスト2",
			Reward:       200,
			ContractName: "無料",
		},
	}

	// DTOを返却する。
	return getTasksUsecaseDto, nil
}

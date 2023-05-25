package task

import "go-ca/internal/app/application/usecase/task"

type TasksQueryService struct {
}

func NewTasksQueryService() *TasksQueryService {
	return &TasksQueryService{}
}

func (t TasksQueryService) GetTasksContainContract() ([]task.GetTasksUsecaseDto, error) {
	// TODO DBアクセスを行いデータを取得
	// 複数の集約を跨ぐ
	getTasksUsecaseDto := []task.GetTasksUsecaseDto{
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

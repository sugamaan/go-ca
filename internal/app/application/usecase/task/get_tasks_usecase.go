package task

import (
	"go-ca/internal/app/domain/task/entity"
	"go-ca/internal/app/domain/task/repository"
	"go-ca/internal/app/domain/task/valueobject"
)

type GetTasksUsecase struct {
	tasksRepository repository.TasksRepository
}

type GetTasksDto struct {
	TaskId uint64
	Name   string
	Reward uint64
}

func NewGetTasksUsecase(tasksRepository repository.TasksRepository) GetTasksUsecase {
	return GetTasksUsecase{tasksRepository: tasksRepository}
}

func (u *GetTasksUsecase) GetTasks() []*entity.Task {
	// DBからタスク一覧を取得
	tasksDataModel, err := u.tasksRepository.GetAllTasks()
	if err != nil {
		// TODO エラー処理
		panic(err)
	}

	// tasksDataModelからDTOへ詰め替え
	var taskDtos []*GetTasksDto
	for _, taskDataModel := range tasksDataModel {
		taskDto := &GetTasksDto{
			TaskId: taskDataModel.TaskId,
			Name:   taskDataModel.Name,
			Reward: taskDataModel.Reward,
		}
		taskDtos = append(taskDtos, taskDto)
	}

	// DTOからNewTaskを使って[]entity.Taskへ詰め替える
	tasks := make([]*entity.Task, len(taskDtos))
	for i, taskDto := range taskDtos {
		reward, err := valueobject.NewReward(taskDto.Reward)
		if err != nil {
			// TODO エラー処理
			panic(err)
		}
		task, err := entity.NewTask(
			taskDto.TaskId,
			taskDto.Name,
			reward)
		if err != nil {
			// TODO エラー処理
			panic(err)
		}
		tasks[i] = task
	}

	// UI層へ受け渡す
	return tasks
}

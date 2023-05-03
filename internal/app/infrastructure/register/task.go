package register

import (
	taskController "go-ca/internal/app/adaptor/controller/task"
	taskUsecase "go-ca/internal/app/application/usecase/task"
	taskRepository "go-ca/internal/app/infrastructure/repository/task"
	"net/http"
)

func Tasks(w http.ResponseWriter, r *http.Request) {
	// DI
	tasksRepository := taskRepository.NewTasksRepository()
	GetTasksUsecase := taskUsecase.NewGetTasksUsecase(&tasksRepository)
	tasksController := taskController.NewTasksController(&GetTasksUsecase)
	tasksController.GetTasks(w, r)
}

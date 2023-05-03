package register

import (
	tasksController "go-ca/internal/app/adaptor/controller/task"
	tasksUsecase "go-ca/internal/app/application/usecase/task"
	tasksRepository "go-ca/internal/app/infrastructure/repository/task"
	"net/http"
)

func Tasks(w http.ResponseWriter, r *http.Request) {
	// DI
	tasksRepository := tasksRepository.NewTasksRepository()
	tasksUsecase := tasksUsecase.NewGetTasksUsecase(&tasksRepository)
	tasksController := tasksController.NewTasksController(&tasksUsecase)
	tasksController.GetTasks(w, r)
}

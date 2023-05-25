package register

import (
	tasksController "go-ca/internal/app/adaptor/controller/task"
	tasksUsecase "go-ca/internal/app/application/usecase/task"
	tasksQueryService "go-ca/internal/app/infrastructure/queryservice/task"
	"net/http"
)

func Tasks(w http.ResponseWriter, r *http.Request) {
	// DI
	tasksQueryService := tasksQueryService.NewTasksQueryService()
	tasksUsecase := tasksUsecase.NewGetTasksUsecase(tasksQueryService)
	tasksController := tasksController.NewTasksController(&tasksUsecase)
	tasksController.GetTasks(w, r)
}

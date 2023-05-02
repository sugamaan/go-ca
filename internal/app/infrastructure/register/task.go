package register

import (
	task3 "go-ca/internal/app/adaptor/controller/task"
	task2 "go-ca/internal/app/application/usecase/task"
	"go-ca/internal/app/infrastructure/repository/task"
	"net/http"
)

func Tasks(w http.ResponseWriter, r *http.Request) {
	// DI
	tasksRepository := task.NewTasksRepository()
	GetTasksUsecase := task2.NewGetTasksUsecase(&tasksRepository)
	tasksController := task3.NewTasksController(&GetTasksUsecase)
	tasksController.GetTasks(w, r)
}

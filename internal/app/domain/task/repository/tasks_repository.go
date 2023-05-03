package repository

import (
	"go-ca/internal/app/domain/task/entity"
)

type TasksRepository interface {
	GetAllTasks() ([]*entity.Task, error)
}

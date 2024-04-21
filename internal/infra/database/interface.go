package database

import (
	"github.com/jdanilocorrea/go-todo-app/internal/entity"
)

type TaskInterface interface {
	CreateTask(task *entity.Task) error
	FindAll(page, limit int, sort string) ([]entity.Task, error)
}

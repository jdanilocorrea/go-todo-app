package database

import (
	"errors"
	"time"

	"github.com/jdanilocorrea/go-todo-app/internal/entity"
	"gorm.io/gorm"
)

type Task struct {
	DB *gorm.DB
}

func NewTask(db *gorm.DB) *Task {
	return &Task{
		DB: db,
	}
}

func (t *Task) CreateTask(task *entity.Task) error {
	if task == nil {
		return errors.New("a tarefa não pode ser nula")
	}
	if task.Name == "" {
		return errors.New("o nome da tarefa não pode estar vazio")
	}
	if task.Description == "" {
		return errors.New("o descrição da tarefa não pode estar vazio")
	}
	if task.Priority == 0 {
		return errors.New("o prioridade da tarefa não pode estar vazio")
	}
	task.CreatedAt = time.Now() // Registra o momento da criação da tarefa.
	task.UpdatedAt = time.Now() // Registra o momento da última atualização da tarefa.

	t.DB.Create(task)
	return nil // Retorna a tarefa criada e nenhum erro.
}

package entity

import (
	"errors"
	"time"

	"github.com/jdanilocorrea/go-todo-app/pkg/entity"
)

// Task representa uma tarefa com identificação, nome, descrição, prioridade e timestamps.
type Task struct {
	ID          entity.ID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask(name, description string, priority int) (*Task, error) {

	if name == "" {
		return nil, errors.New("o nome da tarefa não pode estar vazio")
	}
	if description == "" {
		return nil, errors.New("a descrição da tarefa não pode estar vazia")
	}
	if priority < 1 || priority > 5 { // Supondo que a prioridade deve estar entre 1 e 5
		return nil, errors.New("a prioridade deve estar entre 1 e 5")
	}

	currentTime := time.Now()
	return &Task{
		ID:          entity.NewID(),
		Name:        name,
		Description: description,
		Priority:    priority,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}, nil
}

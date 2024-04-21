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

func (t *Task) FindAll(page, limit int, sort string) ([]entity.Task, error) {
	var tasks []entity.Task

	// Correção da ordenação para garantir que o espaço esteja correto.
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc" // Define um padrão caso o sort esteja vazio ou seja inválido.
	}
	order := "created_at " + sort // Garante que o SQL esteja bem formado.

	// Aplica paginação somente se ambos, page e limit, forem maiores que zero.
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit // Calcula o offset corretamente.
		err := t.DB.Limit(limit).Offset(offset).Order(order).Find(&tasks).Error
		if err != nil {
			return nil, err
		}
	} else {
		// Se a paginação não for aplicada, retorna todas as tarefas com a ordem especificada.
		err := t.DB.Order(order).Find(&tasks).Error
		if err != nil {
			return nil, err
		}
	}

	return tasks, nil
}

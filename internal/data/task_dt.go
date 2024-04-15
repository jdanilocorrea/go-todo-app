package data

import (
	"errors"
	"sync"
	"time"

	"github.com/jdanilocorrea/go-todo-app/internal/entity"
)

var (
	tasks   = make(map[int]*entity.Task) // Mapa armazenando tarefas com seu ID como chave.
	nextID  = 1                          // Variável para gerar um ID único para cada nova tarefa.
	tasksMu sync.Mutex                   // Mutex para controlar o acesso ao mapa de tarefas.
)

// CreateTask adiciona uma nova tarefa ao mapa de tarefas e retorna a tarefa criada.
// Retorna um erro se a tarefa for nula ou se houver algum problema com os dados da tarefa.
func CreateTask(task *entity.Task) (*entity.Task, error) {
	if task == nil {
		return nil, errors.New("a tarefa não pode ser nula")
	}

	if task.Name == "" {
		return nil, errors.New("o nome da tarefa não pode estar vazio")
	}

	tasksMu.Lock() // Garante exclusividade ao modificar o mapa de tarefas.
	defer tasksMu.Unlock()

	// Verifica se o ID já está em uso (não é necessário se nextID é sempre incrementado corretamente)
	if _, exists := tasks[nextID]; exists {
		return nil, errors.New("conflito de ID, já existe uma tarefa com o próximo ID")
	}

	task.ID = nextID            // Atribui um ID único à nova tarefa.
	nextID++                    // Incrementa o ID para uso na próxima tarefa.
	task.CreatedAt = time.Now() // Registra o momento da criação da tarefa.
	task.UpdatedAt = time.Now() // Registra o momento da última atualização da tarefa.
	tasks[task.ID] = task       // Adiciona a tarefa ao mapa.

	return task, nil // Retorna a tarefa criada e nenhum erro.
}

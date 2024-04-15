package data

import (
	"testing"

	"github.com/jdanilocorrea/go-todo-app/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	// Criar uma nova tarefa usando o construtor NewTask
	task, err := entity.NewTask("Comprar", "ir ao mercado", 1)
	assert.Nil(t, err, "erro ao criar a tarefa não deveria ocorrer")
	assert.NotNil(t, task, "a tarefa criada não deveria ser nula")

	// Adicionar a tarefa ao mapa de tarefas
	createdTask, err := CreateTask(task)
	assert.Nil(t, err, "erro ao adicionar a tarefa não deveria ocorrer")
	assert.NotNil(t, createdTask, "a tarefa adicionada não deveria ser nula")

	// Verificações básicas sobre a tarefa retornada
	assert.Equal(t, "Comprar", task.Name, "o nome da tarefa deve ser o esperado")
	assert.Equal(t, "ir ao mercado", task.Description, "a descrição da tarefa deve ser a esperada")
	assert.NotEmpty(t, task.ID, "o ID da tarefa não deve estar vazio")
	assert.True(t, !createdTask.CreatedAt.IsZero(), "CreatedAt deve ser definido")
	assert.True(t, !createdTask.UpdatedAt.IsZero(), "UpdatedAt deve ser definido")

}

package database

import (
	"testing"

	"github.com/jdanilocorrea/go-todo-app/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Db(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Task{})
	return db
}

func TestCreateTask(t *testing.T) {
	// Criar uma nova tarefa usando o construtor NewTask
	task, err := entity.NewTask("Comprar", "ir ao mercado", 1)
	assert.NotNil(t, task, "a tarefa criada não deveria ser nula")
	assert.Nil(t, err, "erro ao criar a tarefa não deveria ocorrer")
	db := Db(t)
	taskDB := NewTask(db)
	err = taskDB.CreateTask(task)
	// Adicionar a tarefa ao mapa de tarefas
	assert.Nil(t, err, "erro ao adicionar a tarefa não deveria ocorrer")
	// Verificações básicas sobre a tarefa retornada
	assert.Equal(t, "Comprar", task.Name, "o nome da tarefa deve ser o esperado")
	assert.Equal(t, "ir ao mercado", task.Description, "a descrição da tarefa deve ser a esperada")
	assert.NotEmpty(t, task.ID, "o ID da tarefa não deve estar vazio")

}

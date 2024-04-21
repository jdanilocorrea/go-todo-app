package database

import (
	"fmt"
	"math/rand"
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
	task, err := entity.NewTask("Task 1", "Description task 1", 1)
	// assert.Nil(t, task.ID)
	assert.NotNil(t, task, "a tarefa criada não deveria ser nula")
	assert.Nil(t, err, "erro ao criar a tarefa não deveria ocorrer")
	db := Db(t)
	taskDB := NewTask(db)
	err = taskDB.CreateTask(task)
	// Adicionar a tarefa ao mapa de tarefas
	assert.Nil(t, err, "erro ao adicionar a tarefa não deveria ocorrer")
	// Verificações básicas sobre a tarefa retornada
	assert.Equal(t, "Task 1", task.Name, "o nome da tarefa deve ser o esperado")
	assert.Equal(t, "Description task 1", task.Description, "a descrição da tarefa deve ser a esperada")
	assert.NotEmpty(t, task.ID, "o ID da tarefa não deve estar vazio")

}
func TestFindAll(t *testing.T) {
	db := Db(t)
	taskDB := NewTask(db)
	for i := 1; i < 24; i++ {
		task_name := fmt.Sprintf("Task %d", i)
		task_description := fmt.Sprintf("Description Task %d", i)
		num := rand.Intn(5) + 1
		task, err := entity.NewTask(task_name, task_description, num)
		assert.NoError(t, err, "Deveria não ter ocorrido erro ao buscar tarefas")
		assert.NotNil(t, task, "O resultado não deveria ser nil")
		// fmt.Printf("ID: %d \n Name: %s\n Description: %s\n Priority: %d\n Created At: %s\n Update At: %s ", task.ID, task.Name, task.Description, task.Priority, task.CreatedAt, task.UpdatedAt)
		err = taskDB.CreateTask(task)
		assert.Nil(t, err)
	}

	assert.NotNil(t, taskDB)
	tasks, err := taskDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.NotNil(t, tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 10)
	assert.Equal(t, "Task 1", tasks[0].Name)
	assert.Equal(t, "Task 10", tasks[9].Name)

	tasks, err = taskDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, tasks, 10)
	assert.Equal(t, "Task 11", tasks[0].Name)
	assert.Equal(t, "Task 20", tasks[9].Name)

	tasks, err = taskDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, tasks, 3)
	assert.Equal(t, "Task 21", tasks[0].Name)
	assert.Equal(t, "Task 23", tasks[2].Name)

}

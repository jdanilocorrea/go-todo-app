package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	task, err := NewTask("Comprar", "ir ao mercado", 1)
	assert.Nil(t, err)
	assert.NotNil(t, task)
	assert.NotEmpty(t, task.Name)
	assert.NotEmpty(t, task.Description)
	assert.Equal(t, "Comprar", task.Name)
	assert.Equal(t, "ir ao mercado", task.Description)
}

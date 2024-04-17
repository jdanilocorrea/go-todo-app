package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jdanilocorrea/go-todo-app/internal/entity"
	"github.com/jdanilocorrea/go-todo-app/internal/infra/database"
	"github.com/jdanilocorrea/go-todo-app/internal/infra/dto"
)

type TaskHandler struct {
	TaskDB database.TaskInterface
}

func NewTaskHandler(db database.TaskInterface) *TaskHandler {

	return &TaskHandler{
		TaskDB: db,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task_input dto.CreateTaskInput
	err := json.NewDecoder(r.Body).Decode(&task_input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t, err := entity.NewTask(task_input.Name, task_input.Description, task_input.Priority)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.TaskDB.CreateTask(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

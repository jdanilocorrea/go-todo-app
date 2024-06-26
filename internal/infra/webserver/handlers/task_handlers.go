package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
		log.Println(err)
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

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	tasks, err := h.TaskDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

package main

import (
	"net/http"

	"github.com/jdanilocorrea/go-todo-app/internal/entity"
	"github.com/jdanilocorrea/go-todo-app/internal/infra/database"
	"github.com/jdanilocorrea/go-todo-app/internal/infra/webserver/handlers"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// configs, err := configs.LoadConfig(".")
	// token := configs.TokenAuth
	// if err != nil {
	// 	panic(err)
	// }
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Task{})
	taskDB := database.NewTask(db)
	taskHandler := handlers.NewTaskHandler(taskDB)

	r := chi.NewRouter()
	// r.Use(middleware.Logger)

	r.Route("/tasks", func(r chi.Router) {
		// r.Use(jwtauth.Verifier(token))
		// r.Use(jwtauth.Authenticator)
		r.Post("/", taskHandler.CreateTask)
		r.Get("/", taskHandler.GetTasks)
		// r.Get("/{id}", taskHandler.GetTask)
		// r.Put("/{id}", taskHandler.UpdateTask)
		// r.Delete("/{id}", taskHandler.DeleteTask)
	})

	http.ListenAndServe(":8080", r)
}

package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"

	"TaskCrud/data/repositories"
	"TaskCrud/handlers"
	"TaskCrud/services"
)

func main() {
	conn, err := pgx.Connect(context.Background(),
		"postgres://admin:admin123@localhost:5432/taskdb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	taskRepository := repositories.NewTaskRepository(conn)
	taskService := services.NewTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)

	http.HandleFunc("/tasks/create", taskHandler.CreateTask)
	http.HandleFunc("/tasks/all", taskHandler.GetAll)
	http.HandleFunc("/tasks/get", taskHandler.GetByID)
	http.HandleFunc("/tasks/update", taskHandler.UpdateTask)
	http.HandleFunc("/tasks/delete", taskHandler.Delete)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

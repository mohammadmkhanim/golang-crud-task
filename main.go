package main

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	"TaskCrud/DTOs/requests"
	"TaskCrud/data/repositories"
	_ "TaskCrud/docs"
	"TaskCrud/handlers"
	"TaskCrud/middlewares"
	"TaskCrud/services"
	"TaskCrud/utils"
)

// @title TaskCrud API
// @version 1.0
func main() {
	utils.LogInfo("main", "starting application")

	var connectionString string = "postgres://admin:admin123@localhost:5432/taskdb?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		panic(err)
	}
	MigrateDatabase(connectionString)

	taskRepository := repositories.NewTaskRepository(conn)

	taskService := services.NewTaskService(taskRepository)

	taskHandler := handlers.NewTaskHandler(taskService)

	// TODO - get id from the path, not query parameter
	http.HandleFunc("/tasks/create", middlewares.RequireMethod(http.MethodPost, middlewares.WithValidatedBody[requests.CreateTaskReq](taskHandler.CreateTask)))
	http.HandleFunc("/tasks/all", middlewares.RequireMethod(http.MethodGet, taskHandler.GetAll))
	http.HandleFunc("/tasks/get", middlewares.RequireMethod(http.MethodGet, taskHandler.GetByID))
	http.HandleFunc("/tasks/update", middlewares.RequireMethod(http.MethodPut, middlewares.WithValidatedBody[requests.UpdateTaskReq](taskHandler.UpdateTask)))
	http.HandleFunc("/tasks/delete", middlewares.RequireMethod(http.MethodDelete, taskHandler.Delete))

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	utils.LogInfo("main", "server running on :8080")
	http.ListenAndServe(":8080", nil)
}

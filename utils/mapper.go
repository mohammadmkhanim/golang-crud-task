package utils

import (
	"TaskCrud/DTOs/requests"
	"TaskCrud/DTOs/responses"
	"TaskCrud/data/models"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func MapCreateTaskReq(req *requests.CreateTaskReq) *models.Task {
	return &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}
}

func MapUpdateTaskReq(req *requests.UpdateTaskReq) *models.Task {
	return &models.Task{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}
}

func ToTaskRes(task *models.Task) *responses.TaskResponse {
	return &responses.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
	}
}

func ToTasksRes(tasks []models.Task) []responses.TaskResponse {
	result := make([]responses.TaskResponse, len(tasks))

	for i, task := range tasks {
		result[i] = *ToTaskRes(&task)
	}

	return result
}

// TODO - is there better soloution than hardcoding?
func MapValidationErrors(err error) []string {
	var errors []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			switch e.Tag() {

			case "required":
				errors = append(errors, fmt.Sprintf("%s is required.", e.Field()))

			case "min":
				errors = append(errors, fmt.Sprintf("%s must be at least %s characters.", e.Field(), e.Param()))

			case "max":
				errors = append(errors, fmt.Sprintf("%s must be at most %s characters.", e.Field(), e.Param()))

			case "oneof":
				errors = append(errors, fmt.Sprintf("%s is invalid.", e.Field()))

			case "uuid":
				errors = append(errors, fmt.Sprintf("%s must be a valid UUID.", e.Field()))

			case "taskstatus":
				errors = append(errors, fmt.Sprintf("%s is invalid.", e.Field()))

			default:
				errors = append(errors, fmt.Sprintf("%s is invalid.", e.Field()))
			}
		}

		return errors
	}

	return []string{err.Error()}
}

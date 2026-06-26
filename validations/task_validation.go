package validations

import (
	"TaskCrud/data/models"
	"errors"
)

func isValidStatus(status models.TaskStatus) bool {
	switch status {
	case models.Todo, models.InProgress, models.Done:
		return true
	default:
		return false
	}
}

func ValidateCreateTask(task *models.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	if !isValidStatus(task.Status) {
		return errors.New("invalid status")
	}
	return nil
}

func ValidateUpdateTask(task *models.Task) error {
	if task.ID == "" {
		return errors.New("id is required")
	}
	if task.Title == "" {
		return errors.New("title is required")
	}
	if !isValidStatus(task.Status) {
		return errors.New("invalid status")
	}
	return nil
}

package requests

import "TaskCrud/data/models"

type CreateTaskReq struct {
	Title       string            `json:"title" validate:"required,min=3,max=100"`
	Description *string           `json:"description,omitempty" validate:"omitempty,max=1000"`
	Status      models.TaskStatus `json:"status" validate:"required,oneof=todo in_progress done"`
}

type UpdateTaskReq struct {
	ID          string            `json:"id" validate:"required,uuid"`
	Title       string            `json:"title" validate:"required,min=3,max=100"`
	Description *string           `json:"description,omitempty" validate:"omitempty,max=1000"`
	Status      models.TaskStatus `json:"status" validate:"required,oneof=todo in_progress done"`
}

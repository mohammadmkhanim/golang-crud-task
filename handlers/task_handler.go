package handlers

import (
	"TaskCrud/DTOs"
	"TaskCrud/DTOs/requests"
	"TaskCrud/data/models"
	"TaskCrud/services"
	"TaskCrud/utils"
	"TaskCrud/validations"
	"encoding/json"
	"net/http"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(s *services.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if !CheckHttpMethod(w, r, http.MethodPost) {
		return
	}

	var req requests.CreateTaskReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		DTOs.Error(
			w,
			http.StatusBadRequest,
			"Invalid request body",
		)
		return
	}

	if err := validations.Validate.Struct(req); err != nil {
		DTOs.Error(
			w,
			http.StatusBadRequest,
			"Validation failed",
			utils.MapValidationErrors(err)...,
		)
		return
	}

	task := utils.MapCreateTaskReq(&req)

	err = h.service.CreateTask(r.Context(), task)
	if err != nil {
		DTOs.Error(
			w,
			http.StatusInternalServerError,
			"Failed to create task",
			err.Error(),
		)
		return
	}

	res := utils.ToTaskRes(task)

	DTOs.Success(
		w,
		http.StatusCreated,
		"Task created successfully",
		res,
	)
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if !CheckHttpMethod(w, r, http.MethodGet) {
		return
	}

	tasks, err := h.service.GetAll(r.Context())
	if err != nil {
		DTOs.Error(
			w,
			http.StatusInternalServerError,
			"Failed to retrieve tasks",
			err.Error(),
		)
		return
	}

	res := utils.ToTasksRes(tasks)

	DTOs.Success(
		w,
		http.StatusOK,
		"Tasks retrieved successfully",
		&res,
	)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if !CheckHttpMethod(w, r, http.MethodGet) {
		return
	}

	id := r.URL.Query().Get("id")

	task, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		DTOs.Error(
			w,
			http.StatusNotFound,
			"Task not found",
		)
		return
	}

	res := utils.ToTaskRes(task)

	DTOs.Success(
		w,
		http.StatusOK,
		"Task retrieved successfully",
		res,
	)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	if !CheckHttpMethod(w, r, http.MethodPut) {
		return
	}

	var req requests.UpdateTaskReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		DTOs.Error(
			w,
			http.StatusBadRequest,
			"Invalid request body",
		)
		return
	}

	if err := validations.Validate.Struct(req); err != nil {
		DTOs.Error(
			w,
			http.StatusBadRequest,
			"Validation failed",
			utils.MapValidationErrors(err)...,
		)
		return
	}

	task := utils.MapUpdateTaskReq(&req)

	existingTask, ok := h.CheckExistTask(w, r, task.ID)
	if !ok {
		return
	}

	err = h.service.UpdateTask(r.Context(), existingTask, task)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	res := utils.ToTaskRes(task)

	DTOs.Success(
		w,
		http.StatusOK,
		"Task updated successfully",
		res,
	)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if !CheckHttpMethod(w, r, http.MethodDelete) {
		return
	}

	id := r.URL.Query().Get("id")

	_, ok := h.CheckExistTask(w, r, id)
	if !ok {
		return
	}

	err := h.service.DeleteTask(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	DTOs.Success[any](
		w,
		http.StatusOK,
		"Task deleted successfully",
		nil,
	)
}

func (h *TaskHandler) CheckExistTask(w http.ResponseWriter, r *http.Request, id string) (*models.Task, bool) {
	existingTask, err := h.service.GetByID(r.Context(), id)
	if existingTask == nil || err != nil {
		DTOs.Error(
			w,
			http.StatusNotFound,
			"Task not found",
		)
		return nil, false
	}
	return existingTask, true
}

func CheckHttpMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		DTOs.Error(
			w,
			http.StatusMethodNotAllowed,
			"Method not allowed",
		)
		return false
	}
	return true
}

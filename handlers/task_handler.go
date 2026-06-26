package handlers

import (
	"TaskCrud/data/models"
	"TaskCrud/services"
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

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := validations.ValidateCreateTask(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateTask(r.Context(), &task)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if !CheckHttpMethod(w, r, http.MethodGet) {
		return
	}

	tasks, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if !CheckHttpMethod(w, r, http.MethodGet) {
		return
	}

	id := r.URL.Query().Get("id")

	task, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to find the task", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	if !CheckHttpMethod(w, r, http.MethodPut) {
		return
	}

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := validations.ValidateUpdateTask(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existingTask, ok := h.CheckExistTask(w, r, task.ID)
	if !ok {
		return
	}

	err = h.service.UpdateTask(r.Context(), existingTask, &task)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
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

	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) CheckExistTask(w http.ResponseWriter, r *http.Request, id string) (*models.Task, bool) {
	existingTask, err := h.service.GetByID(r.Context(), id)
	if existingTask == nil || err != nil {
		http.Error(w, "Failed to find the task", http.StatusNotFound)
		return nil, false
	}
	return existingTask, true
}

func CheckHttpMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

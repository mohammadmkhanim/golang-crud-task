package handlers

import (
	"TaskCrud/DTOs"
	"TaskCrud/DTOs/requests"
	"TaskCrud/DTOs/responses"
	"TaskCrud/data/models"
	"TaskCrud/services"
	"TaskCrud/utils"
	"net/http"
)

const (
	defaultPage     = 1
	defaultPageSize = 10
	maxPageSize     = 100
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(s *services.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

// @Tags tasks
// @Accept json
// @Produce json
// @Param request body requests.CreateTaskReq true "Task payload"
// @Success 201 {object} DTOs.BaseRes[responses.TaskResponse]
// @Failure 400 {object} DTOs.BaseRes[any]
// @Failure 500 {object} DTOs.BaseRes[any]
// @Router /tasks/create [post]
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request, req requests.CreateTaskReq) {
	utils.LogInfo("CreateTask", "handling create task request")

	task := utils.MapCreateTaskReq(&req)

	err := h.service.CreateTask(r.Context(), task)
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

	utils.LogSuccess("CreateTask", "task created successfully with id {0}", task.ID)
}

// @Tags tasks
// @Produce json
// @Param status query string false "Filter by status" Enums(todo, in_progress, done)
// @Param order query string false "Sort order" Enums(asc, desc)
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size (max 100)"
// @Success 200 {object} DTOs.BaseRes[responses.PaginatedResponse[responses.TaskResponse]]
// @Failure 400 {object} DTOs.BaseRes[any]
// @Failure 500 {object} DTOs.BaseRes[any]
// @Router /tasks/all [get]
func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("GetAll", "handling get all tasks request")

	query, errMsg := parseTaskListQuery(r)
	if errMsg != "" {
		DTOs.Error(
			w,
			http.StatusBadRequest,
			errMsg,
		)
		return
	}

	tasks, totalItems, err := h.service.GetAll(r.Context(), query.Status, query.Order, query.Page, query.PageSize)
	if err != nil {
		DTOs.Error(
			w,
			http.StatusInternalServerError,
			"Failed to retrieve tasks",
			err.Error(),
		)
		return
	}

	res := responses.PaginatedResponse[responses.TaskResponse]{
		Items: utils.ToTasksRes(tasks),
		Pagination: responses.PaginationMeta{
			Page:       query.Page,
			PageSize:   query.PageSize,
			TotalItems: totalItems,
			TotalPages: utils.TotalPages(totalItems, query.PageSize),
		},
	}

	DTOs.Success(
		w,
		http.StatusOK,
		"Tasks retrieved successfully",
		&res,
	)

	utils.LogSuccess("GetAll", "retrieved {0} of {1} tasks successfully (page {2})", len(tasks), totalItems, query.Page)
}

// @Tags tasks
// @Produce json
// @Param id query string true "Task ID"
// @Success 200 {object} DTOs.BaseRes[responses.TaskResponse]
// @Failure 404 {object} DTOs.BaseRes[any]
// @Router /tasks/get [get]
func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("GetByID", "handling get task by id request")

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

	utils.LogSuccess("GetByID", "task {0} retrieved successfully", id)
}

// @Tags tasks
// @Accept json
// @Produce json
// @Param request body requests.UpdateTaskReq true "Task payload"
// @Success 200 {object} DTOs.BaseRes[responses.TaskResponse]
// @Failure 400 {object} DTOs.BaseRes[any]
// @Failure 404 {object} DTOs.BaseRes[any]
// @Failure 500 {object} DTOs.BaseRes[any]
// @Router /tasks/update [put]
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request, req requests.UpdateTaskReq) {
	utils.LogInfo("UpdateTask", "handling update task request")

	task := utils.MapUpdateTaskReq(&req)

	existingTask, ok := h.checkExistTask(w, r, task.ID)
	if !ok {
		return
	}

	err := h.service.UpdateTask(r.Context(), existingTask, task)
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

	utils.LogSuccess("UpdateTask", "task {0} updated successfully", task.ID)
}

// @Tags tasks
// @Produce json
// @Param id query string true "Task ID"
// @Success 200 {object} DTOs.BaseRes[any]
// @Failure 404 {object} DTOs.BaseRes[any]
// @Failure 500 {object} DTOs.BaseRes[any]
// @Router /tasks/delete [delete]
func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("Delete", "handling delete task request")

	id := r.URL.Query().Get("id")

	_, ok := h.checkExistTask(w, r, id)
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

	utils.LogSuccess("Delete", "task {0} deleted successfully", id)
}

func (h *TaskHandler) checkExistTask(w http.ResponseWriter, r *http.Request, id string) (*models.Task, bool) {
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

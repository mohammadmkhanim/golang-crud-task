package handlers

import (
	"net/http"

	"TaskCrud/data/models"
	"TaskCrud/utils"
)

type taskListQuery struct {
	Status   *models.TaskStatus
	Order    models.SortOrder
	Page     int
	PageSize int
}

func parseTaskListQuery(r *http.Request) (*taskListQuery, string) {
	var statusFilter *models.TaskStatus
	if statusParam := r.URL.Query().Get("status"); statusParam != "" {
		status := models.TaskStatus(statusParam)
		if !status.IsValidStatus() {
			return nil, "Invalid status filter"
		}
		statusFilter = &status
	}

	order := models.Asc
	if orderParam := r.URL.Query().Get("order"); orderParam != "" {
		order = models.SortOrder(orderParam)
		if !order.IsValidSortOrder() {
			return nil, "Invalid order parameter, must be 'asc' or 'desc'"
		}
	}

	pagination, errMsg := utils.ParsePagination(r, defaultPage, defaultPageSize, maxPageSize)
	if errMsg != "" {
		return nil, errMsg
	}

	return &taskListQuery{
		Status:   statusFilter,
		Order:    order,
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
	}, ""
}

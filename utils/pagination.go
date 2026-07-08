package utils

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Page     int
	PageSize int
}

func ParsePagination(r *http.Request, defaultPage, defaultPageSize, maxPageSize int) (Pagination, string) {
	page := defaultPage
	if pageParam := r.URL.Query().Get("page"); pageParam != "" {
		parsedPage, err := strconv.Atoi(pageParam)
		if err != nil || parsedPage < 1 {
			return Pagination{}, "Invalid page parameter, must be a positive integer"
		}
		page = parsedPage
	}

	pageSize := defaultPageSize
	if pageSizeParam := r.URL.Query().Get("pageSize"); pageSizeParam != "" {
		parsedPageSize, err := strconv.Atoi(pageSizeParam)
		if err != nil || parsedPageSize < 1 || parsedPageSize > maxPageSize {
			return Pagination{}, "Invalid pageSize parameter, must be an integer between 1 and 100"
		}
		pageSize = parsedPageSize
	}

	return Pagination{Page: page, PageSize: pageSize}, ""
}

func TotalPages(totalItems int64, pageSize int) int {
	return int((totalItems + int64(pageSize) - 1) / int64(pageSize))
}

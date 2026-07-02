package responses

type PaginationMeta struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int   `json:"totalPages"`
}

type PaginatedResponse[T any] struct {
	Items      []T            `json:"items"`
	Pagination PaginationMeta `json:"pagination"`
}

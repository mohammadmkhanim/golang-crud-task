package DTOs

import (
	"encoding/json"
	"net/http"
)

type BaseRes[T any] struct {
	Success    bool     `json:"success"`
	StatusCode int      `json:"statusCode"`
	Message    string   `json:"message,omitempty"`
	Data       *T       `json:"data,omitempty"`
	Errors     []string `json:"errors,omitempty"`
}

func Success[T any](w http.ResponseWriter, statusCode int, message string, data *T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(BaseRes[T]{
		Success:    true,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}

func Error(w http.ResponseWriter, statusCode int, message string, errors ...string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(BaseRes[any]{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
		Errors:     errors,
	})
}

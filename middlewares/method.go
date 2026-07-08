package middlewares

import (
	"net/http"

	"TaskCrud/DTOs"
)

func RequireMethod(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			DTOs.Error(
				w,
				http.StatusMethodNotAllowed,
				"Method not allowed",
			)
			return
		}
		next(w, r)
	}
}

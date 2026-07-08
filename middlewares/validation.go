package middlewares

import (
	"encoding/json"
	"net/http"

	"TaskCrud/DTOs"
	"TaskCrud/utils"
	"TaskCrud/validations"
)

func WithValidatedBody[T any](next func(w http.ResponseWriter, r *http.Request, body T)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body T
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			DTOs.Error(
				w,
				http.StatusBadRequest,
				"Invalid request body",
			)
			return
		}

		if err := validations.Validate.Struct(body); err != nil {
			DTOs.Error(
				w,
				http.StatusBadRequest,
				"Validation failed",
				utils.MapValidationErrors(err)...,
			)
			return
		}

		next(w, r, body)
	}
}

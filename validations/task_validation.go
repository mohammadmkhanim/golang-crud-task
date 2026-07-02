package validations

import (
	"TaskCrud/data/models"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func init() {
	Validate.RegisterValidation("taskstatus", validateTaskStatus)
}

func validateTaskStatus(fl validator.FieldLevel) bool {
	status, ok := fl.Field().Interface().(models.TaskStatus)
	if !ok {
		return false
	}

	return status.IsValidStatus()
}

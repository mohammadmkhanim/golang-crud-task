package utils

import "github.com/google/uuid"

func NewDbId() string {
	return uuid.New().String()
}

package utils

import (
	"github.com/satori/go.uuid"
)

func GenGUID() string {
	return uuid.NewV4().String()
}

package utils

import (
	"github.com/satori/go.uuid"
)

// GenerateUUID generate an uuid
func GenerateUUID() string {
	u := uuid.NewV4()
	return u.String()

}

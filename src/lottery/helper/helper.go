package helper

import (
	"github.com/satori/go.uuid"
)

func UUIDGen() uuid.UUID {
	return uuid.NewV4()
}

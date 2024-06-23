package types

import (
	"github.com/google/uuid"
)

type ID uuid.UUID

func NewID() uuid.UUID {
	return uuid.New()
}

func ValidateID(newID string) (ID, error) {
	id, err := uuid.Parse(newID)
	return ID(id), err
}

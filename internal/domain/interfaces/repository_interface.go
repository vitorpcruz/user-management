package interfaces

import (
	"github.com/google/uuid"
)

type RepositoryInterface[TEntity interface{}] interface {
	GetByID(ID uuid.UUID) (*TEntity, error)
	// GetAll(page, rows int) []TEntity // TODO: Implement pagination
	// Create(entity TEntity)
	// Delete(entity TEntity)
	// Update(entity TEntity)
}

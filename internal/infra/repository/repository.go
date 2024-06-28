package repository

import (
	"github.com/google/uuid"
	"github.com/vitorpcruz/sso/internal/domain/util"
	"gorm.io/gorm"
)

type Repository[TEntity interface{}] struct {
	db        *gorm.DB
	tableName string
}

func NewBasicRepo[TEntity interface{}](
	db *gorm.DB,
	entity TEntity,
	tableName string,
) *Repository[TEntity] {
	return &Repository[TEntity]{
		db:        db,
		tableName: util.ConvertToMySqlNamePattern(tableName),
	}
}

func (repo *Repository[TEntity]) GetByID(ID uuid.UUID) (*TEntity, error) {
	var entity TEntity
	err := repo.db.First(&entity, ID).Error
	if err != nil {
		return nil, err
	}

	return &entity, err
}

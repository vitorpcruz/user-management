package repository

import (
	"github.com/vitorpcruz/sso/internal/domain/entity"
	"github.com/vitorpcruz/sso/internal/domain/interfaces"
	"gorm.io/gorm"
)

const (
	tableName = "customer"
)

type CustomerRepository struct {
	interfaces.RepositoryInterface[entity.Customer]
	orm *gorm.DB
}

func NewCustomerRepository(
	orm *gorm.DB,
	repository interfaces.RepositoryInterface[entity.Customer],
) *CustomerRepository {
	return &CustomerRepository{
		orm:                 orm,
		RepositoryInterface: repository,
	}
}

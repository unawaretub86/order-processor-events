package repository

import (
	"github.com/unawaretub86/order-processor-events/internal/domain/entities"
	"github.com/unawaretub86/order-processor-events/internal/domain/repository/database"
)

type (
	RepositoryOrder interface {
		CreateOrder(*entities.OrderRequest, string) (*string, error)
	}

	repositoryOrder struct {
		database database.Database
	}
)

func NewRepository(database database.Database) RepositoryOrder {
	return &repositoryOrder{
		database: database,
	}
}

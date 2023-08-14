package repository

import "github.com/unawaretub86/order-processor-events/internal/domain/repository/database"

type (
	RepositoryOrder interface {
		CreateOrder(body, requestId string) error
	}

	repositoryOrder struct {
		database database.Database
	}
)

func NewDataBase(database database.Database) RepositoryOrder {
	return &repositoryOrder{
		database: database,
	}
}

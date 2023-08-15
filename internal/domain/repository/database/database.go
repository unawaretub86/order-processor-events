package database

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/unawaretub86/order-processor-events/internal/domain/entities"
)

type (
	Database interface {
		CreateOrder(*entities.OrderRequest, string) (*string, error)
	}

	databaseOrder struct {
		dynamodb dynamodbiface.DynamoDBAPI
		table    string
	}
)

func NewDataBase(dynamodbClient dynamodbiface.DynamoDBAPI) Database {
	const tableName = "orders"

	return &databaseOrder{
		dynamodb: dynamodbClient,
		table:    tableName,
	}
}

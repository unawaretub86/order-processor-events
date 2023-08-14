package database

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type (
	Database interface {
		CreateOrder(body, requestId string) error
	}

	databaseOrder struct {
		dynamodb dynamodbiface.DynamoDBAPI
	}
)

func NewDataBase(dynamodbClient dynamodbiface.DynamoDBAPI) Database {
	return &databaseOrder{
		dynamodb: dynamodbClient,
	}
}

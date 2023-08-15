package database

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"

	"github.com/unawaretub86/order-processor-events/internal/domain/entities"
)

func (d *databaseOrder) CreateOrder(body *entities.OrderRequest, requestId string) (*string, error) {
	orderId := uuid.New().String()

	item := map[string]*dynamodb.AttributeValue{
		"orderId":    {S: aws.String(orderId)},
		"Item":       {S: aws.String(body.Item)},
		"UserId":     {S: aws.String(body.UserID)},
		"Quantity":   {N: aws.String(strconv.Itoa(body.Quantity))},
		"TotalPrice": {N: aws.String(strconv.FormatInt(body.TotalPrice, 10))},
	}

	_, err := d.dynamodb.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(d.table),
		Item:      item,
	})

	fmt.Printf("[RequestId: %s], [PutItem result: %v]", orderId, requestId)

	if err != nil {
		fmt.Printf("[RequestId: %s], [Error: %v]", err, requestId)
		return nil, err
	}

	return &orderId, nil
}

package usecase_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"

	"github.com/unawaretub86/order-processor-events/internal/domain/entities"
	"github.com/unawaretub86/order-processor-events/internal/domain/repository/mocks"
	"github.com/unawaretub86/order-processor-events/internal/domain/usecase"
)

func TestCreateOrder_Success(t *testing.T) {
	mockRepo := &mocks.Mocks{}
	useCase := usecase.NewUseOrder(mockRepo)

	validBody := `{
    "user_id": "1234pruebaCompleta",
	"item": "pruebaCompleta",
	"quantity": 1111,
	"total_price": 2222
}`
	requestID := "1234567890"

	mockRepo.CreateOrderFunc = func(order *entities.OrderRequest, requestId string) (*string, error) {
		userID := "user_id"
		return &userID, nil
	}

	_, err := useCase.CreateOrder(validBody, requestID)

	mockSQS := mocks.NewMockSQS("us-east-2")
	queueURL := "https://queue.amazonaws.com/80398EXAMPLE/MyQueue"

	messageAttributes := map[string]*sqs.MessageAttributeValue{
		"Source": {
			DataType:    aws.String("String"),
			StringValue: aws.String("order-processor-events"),
		},
	}

	_, err = mockSQS.SendMessage(&sqs.SendMessageInput{
		MessageBody:       aws.String(`{"order_id": "1234567890"}`),
		QueueUrl:          &queueURL,
		MessageAttributes: messageAttributes,
	})
	if err != nil {
		t.Errorf("Error sending message: %v", err)
	}

	if err != nil {
		t.Errorf("Error updating payment: %v", err)
	}

	assert.NoError(t, err)
}

func TestCreatePayment_Error(t *testing.T) {
	mockRepo := &mocks.Mocks{}
	useCase := usecase.NewUseOrder(mockRepo)

	invalidBody := `{invalid_field: 1234567890}`
	requestID := "1234567890"

	expectedError := fmt.Errorf("invalid input data")

	mockRepo.CreateOrderFunc = func(order *entities.OrderRequest, requestId string) (*string, error) {
		userID := "user_id"
		return &userID, nil
	}

	result, err := useCase.CreateOrder(invalidBody, requestID)

	assert.Error(t, err, expectedError)
	assert.Nil(t, result)
}

func TestCreateOrder_CreateError(t *testing.T) {
	mockRepo := &mocks.Mocks{}
	useCase := usecase.NewUseOrder(mockRepo)

	validBody := `{"order_data": "valid_data"}`
	requestID := "1234567890"

	expectedError := fmt.Errorf("error creating order")

	mockRepo.CreateOrderFunc = func(order *entities.OrderRequest, requestId string) (*string, error) {
		return nil, expectedError
	}

	orderID, err := useCase.CreateOrder(validBody, requestID)

	assert.Error(t, err, expectedError)
	assert.Nil(t, orderID)
}

func TestUpdateOrder_Success(t *testing.T) {
	mockRepo := &mocks.Mocks{}
	useCase := usecase.NewUseOrder(mockRepo)

	orderID := "1234567890"
	requestID := "9876543210"

	mockRepo.UpdateOrderFunc = func(orderID, requestId string) error {
		return nil
	}

	err := useCase.UpdateOrder(orderID, requestID)

	// Assert
	assert.NoError(t, err)
}

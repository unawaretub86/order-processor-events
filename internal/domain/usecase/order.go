package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/unawaretub86/order-processor-events/internal/domain/entities"
)

func (useCase useCase) CreateOrder(body, requestId string) (*string, error) {
	order, err := convertToStruct(body, requestId)
	if err != nil {
		fmt.Printf("[RequestId: %s], [Error: %v]", err, requestId)
		return nil, err
	}

	orderId, err := useCase.repositoryOrder.CreateOrder(order, requestId)
	if err != nil {
		fmt.Printf("[RequestId: %s], [Error: %v]", err, requestId)
		return nil, err
	}

	useCase.sendSQS(*&order.TotalPrice, *orderId, requestId)

	return orderId, nil
}

func convertToStruct(body, requestId string) (*entities.OrderRequest, error) {
	var orderRequest entities.OrderRequest
	err := json.Unmarshal([]byte(body), &orderRequest)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error marshaling API Gateway request: %v]", err, requestId)
		return nil, err
	}

	return &orderRequest, nil
}

func (useCase useCase) sendSQS(totalPrice int64, orderID, requestId string) error {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	sqsClient := sqs.New(sess)
	queueURL := useCase.queueURL

	orderEvent := entities.OrderEvent{
		OrderID:    orderID,
		TotalPrice: totalPrice,
	}

	orderJSON, err := json.Marshal(orderEvent)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error marshaling order request: %v]", err, requestId)
		return err
	}

	_, err = sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(orderJSON)),
		QueueUrl:    &queueURL,
	})

	if err != nil {
		fmt.Printf("[RequestId: %s][Error sending message to SQS: %v]", err, requestId)
		return err
	}

	return nil
}

package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/unawaretub86/order-processor-events/internal/domain/usecase"
)

func HandleSQSMessage(ctx context.Context, sqsEvent events.SQSEvent) error {
	lc, _ := lambdacontext.FromContext(ctx)

	requestId := lc.AwsRequestID

	var messageBody string

	for _, record := range sqsEvent.Records {
		messageBody = record.Body
	}

	useHandler := usecase.NewUseOrder()

	return useHandler.CreateOrder(messageBody, requestId)
}

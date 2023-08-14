package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/unawaretub86/order-processor-events/internal/domain/handler"
)

func main() {
	lambda.Start(handler.HandleSQSMessage)
}

package usecase

import (
	"os"

	"github.com/unawaretub86/order-processor-events/internal/domain/repository"
)

type (
	UseCase interface {
		CreateOrder(string, string) (*string, error)
	}

	useCase struct {
		repositoryOrder repository.RepositoryOrder
		queueURL        string
	}
)

func NewUseOrder(repositoryOrder repository.RepositoryOrder) UseCase {
	queueURL := os.Getenv("SQS_URL")

	return &useCase{
		repositoryOrder: repositoryOrder,
		queueURL:        queueURL,
	}
}

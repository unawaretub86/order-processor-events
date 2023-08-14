package usecase

import "github.com/unawaretub86/order-processor-events/internal/domain/repository"

type (
	UseCase interface {
		CreateOrder(body string, requestId string) error
	}

	useCase struct {
		repositoryOrder repository.RepositoryOrder
	}
)

func NewUseOrder() UseCase {
	return &useCase{}
}

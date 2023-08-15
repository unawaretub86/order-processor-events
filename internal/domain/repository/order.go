package repository

import "github.com/unawaretub86/order-processor-events/internal/domain/entities"

func (repositoryOrder repositoryOrder) CreateOrder(body *entities.OrderRequest, requestId string) (*string, error) {
	return repositoryOrder.database.CreateOrder(body, requestId)
}

func (repositoryOrder repositoryOrder) UpdateOrder(orderId, requestId string) error {
	return repositoryOrder.database.UpdateOrder(orderId, requestId)
}

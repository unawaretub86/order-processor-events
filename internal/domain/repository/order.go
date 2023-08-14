package repository

func (repositoryOrder repositoryOrder) CreateOrder(body, requestId string) error {
	return repositoryOrder.database.CreateOrder(body, requestId)
}

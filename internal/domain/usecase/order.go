package usecase

func (useCase useCase) CreateOrder(body, requestId string) error {
	return useCase.repositoryOrder.CreateOrder(body, requestId)
}

package services

import (
	"context"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/ports"
)

type CustomerService struct {
	repository ports.CustomerRepository
}

func NewCustomerService(repository ports.CustomerRepository) *CustomerService {
	return &CustomerService{
		repository: repository,
	}
}

func (service *CustomerService) CreateCustomer(ctx context.Context, customer domain.Customer) (uint, error) {
	customerId, err := service.repository.CreateCustomer(ctx, customer)

	if err != nil {
		return 0, err
	}

	return customerId, nil
}

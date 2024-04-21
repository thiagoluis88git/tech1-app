package services

import (
	"context"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/ports"
	"thiagoluis88git/tech1/pkg/responses"
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
		return 0, responses.GetResponseError(err, "CustomerService")
	}

	return customerId, nil
}

func (service *CustomerService) GetCustomerByCPF(ctx context.Context, cpf string) (domain.Customer, error) {
	customer, err := service.repository.GetCustomerByCPF(ctx, cpf)

	if err != nil {
		return domain.Customer{}, responses.GetResponseError(err, "CustomerService")
	}

	return customer, nil
}

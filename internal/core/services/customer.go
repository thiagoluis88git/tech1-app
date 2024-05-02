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

func (service *CustomerService) CreateCustomer(ctx context.Context, customer domain.Customer) (domain.CustomerResponse, error) {
	customerId, err := service.repository.CreateCustomer(ctx, customer)

	if err != nil {
		return domain.CustomerResponse{}, responses.GetResponseError(err, "CustomerService")
	}

	return domain.CustomerResponse{
		Id: customerId,
	}, nil
}

func (service *CustomerService) UpdateCustomer(ctx context.Context, customer domain.Customer) error {
	err := service.repository.UpdateCustomer(ctx, customer)

	if err != nil {
		return responses.GetResponseError(err, "CustomerService")
	}

	return nil
}

func (service *CustomerService) GetCustomerByCPF(ctx context.Context, cpf string) (domain.Customer, error) {
	customer, err := service.repository.GetCustomerByCPF(ctx, cpf)

	if err != nil {
		return domain.Customer{}, responses.GetResponseError(err, "CustomerService")
	}

	return customer, nil
}

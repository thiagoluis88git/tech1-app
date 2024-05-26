package services

import (
	"context"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type CustomerService struct {
	validateCPFUseCase *ValidateCPFUseCase
	repository         ports.CustomerRepository
}

func NewCustomerService(validateCPFUseCase *ValidateCPFUseCase, repository ports.CustomerRepository) *CustomerService {
	return &CustomerService{
		validateCPFUseCase: validateCPFUseCase,
		repository:         repository,
	}
}

func (service *CustomerService) CreateCustomer(ctx context.Context, customer domain.Customer) (domain.CustomerResponse, error) {
	cleanedCPF, validate := service.validateCPFUseCase.Execute(customer.CPF)

	if !validate {
		return domain.CustomerResponse{}, &responses.BusinessResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid CPF",
		}
	}

	customer.CPF = cleanedCPF
	customerId, err := service.repository.CreateCustomer(ctx, customer)

	if err != nil {
		return domain.CustomerResponse{}, responses.GetResponseError(err, "CustomerService")
	}

	return domain.CustomerResponse{
		Id: customerId,
	}, nil
}

func (service *CustomerService) UpdateCustomer(ctx context.Context, customer domain.Customer) error {
	cleanedCPF, validate := service.validateCPFUseCase.Execute(customer.CPF)

	if !validate {
		return &responses.BusinessResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid CPF",
		}
	}

	customer.CPF = cleanedCPF
	err := service.repository.UpdateCustomer(ctx, customer)

	if err != nil {
		return responses.GetResponseError(err, "CustomerService")
	}

	return nil
}

func (service *CustomerService) GetCustomerById(ctx context.Context, id uint) (domain.Customer, error) {
	customer, err := service.repository.GetCustomerById(ctx, id)

	if err != nil {
		return domain.Customer{}, responses.GetResponseError(err, "CustomerService")
	}

	return customer, nil
}

func (service *CustomerService) GetCustomerByCPF(ctx context.Context, cpf string) (domain.Customer, error) {
	cleanedCPF, validate := service.validateCPFUseCase.Execute(cpf)

	if !validate {
		return domain.Customer{}, &responses.BusinessResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid CPF",
		}
	}

	customer, err := service.repository.GetCustomerByCPF(ctx, cleanedCPF)

	if err != nil {
		return domain.Customer{}, responses.GetResponseError(err, "CustomerService")
	}

	return customer, nil
}

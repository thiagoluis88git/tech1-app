package usecases

import (
	"context"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type CreateCustomerUseCase struct {
	validateCPFUseCase *ValidateCPFUseCase
	repository         ports.CustomerRepository
}

type UpdateCustomerUseCase struct {
	validateCPFUseCase *ValidateCPFUseCase
	repository         ports.CustomerRepository
}

type GetCustomerByCPFUseCase struct {
	validateCPFUseCase *ValidateCPFUseCase
	repository         ports.CustomerRepository
}

type GetCustomerByIdUseCase struct {
	repository ports.CustomerRepository
}

func NewUpdateCustomerUseCase(validateCPFUseCase *ValidateCPFUseCase, repository ports.CustomerRepository) *UpdateCustomerUseCase {
	return &UpdateCustomerUseCase{
		validateCPFUseCase: validateCPFUseCase,
		repository:         repository,
	}
}

func NewCreateCustomerUseCase(validateCPFUseCase *ValidateCPFUseCase, repository ports.CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		validateCPFUseCase: validateCPFUseCase,
		repository:         repository,
	}
}

func NewGetCustomerByCPFUseCase(validateCPFUseCase *ValidateCPFUseCase, repository ports.CustomerRepository) *GetCustomerByCPFUseCase {
	return &GetCustomerByCPFUseCase{
		validateCPFUseCase: validateCPFUseCase,
		repository:         repository,
	}
}

func NewGetCustomerByIdUseCase(repository ports.CustomerRepository) *GetCustomerByIdUseCase {
	return &GetCustomerByIdUseCase{
		repository: repository,
	}
}

func (service *CreateCustomerUseCase) Execute(ctx context.Context, customer domain.Customer) (domain.CustomerResponse, error) {
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

func (service *UpdateCustomerUseCase) Execute(ctx context.Context, customer domain.Customer) error {
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

func (service *GetCustomerByIdUseCase) Execute(ctx context.Context, id uint) (domain.Customer, error) {
	customer, err := service.repository.GetCustomerById(ctx, id)

	if err != nil {
		return domain.Customer{}, responses.GetResponseError(err, "CustomerService")
	}

	return customer, nil
}

func (service *GetCustomerByCPFUseCase) Execute(ctx context.Context, cpf string) (domain.Customer, error) {
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

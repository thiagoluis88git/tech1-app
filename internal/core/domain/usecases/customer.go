package usecases

import (
	"context"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/domain/repository"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type CreateCustomerUseCase struct {
	validateCPFUseCase *ValidateCPFUseCase
	repository         repository.CustomerRepository
}

type UpdateCustomerUseCase struct {
	validateCPFUseCase *ValidateCPFUseCase
	repository         repository.CustomerRepository
}

type GetCustomerByCPFUseCase struct {
	validateCPFUseCase *ValidateCPFUseCase
	repository         repository.CustomerRepository
}

type GetCustomerByIdUseCase struct {
	repository repository.CustomerRepository
}

type LoginCustomerUseCase struct {
	repository repository.CustomerRepository
}

type LoginUnknownCustomerUseCase struct {
	repository repository.CustomerRepository
}

func NewUpdateCustomerUseCase(validateCPFUseCase *ValidateCPFUseCase, repository repository.CustomerRepository) *UpdateCustomerUseCase {
	return &UpdateCustomerUseCase{
		validateCPFUseCase: validateCPFUseCase,
		repository:         repository,
	}
}

func NewCreateCustomerUseCase(validateCPFUseCase *ValidateCPFUseCase, repository repository.CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		validateCPFUseCase: validateCPFUseCase,
		repository:         repository,
	}
}

func NewGetCustomerByCPFUseCase(validateCPFUseCase *ValidateCPFUseCase, repository repository.CustomerRepository) *GetCustomerByCPFUseCase {
	return &GetCustomerByCPFUseCase{
		validateCPFUseCase: validateCPFUseCase,
		repository:         repository,
	}
}

func NewGetCustomerByIdUseCase(repository repository.CustomerRepository) *GetCustomerByIdUseCase {
	return &GetCustomerByIdUseCase{
		repository: repository,
	}
}

func NewLoginCustomerUseCase(repository repository.CustomerRepository) *LoginCustomerUseCase {
	return &LoginCustomerUseCase{
		repository: repository,
	}
}

func NewLoginUnknownCustomerUseCase(repository repository.CustomerRepository) *LoginUnknownCustomerUseCase {
	return &LoginUnknownCustomerUseCase{
		repository: repository,
	}
}

func (service *CreateCustomerUseCase) Execute(ctx context.Context, customer dto.Customer) (dto.CustomerResponse, error) {
	cleanedCPF, validate := service.validateCPFUseCase.Execute(customer.CPF)

	if !validate {
		return dto.CustomerResponse{}, &responses.BusinessResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid CPF",
		}
	}

	customer.CPF = cleanedCPF
	customerId, err := service.repository.CreateCustomer(ctx, customer)

	if err != nil {
		return dto.CustomerResponse{}, responses.GetResponseError(err, "CustomerService")
	}

	return dto.CustomerResponse{
		Id: customerId,
	}, nil
}

func (service *UpdateCustomerUseCase) Execute(ctx context.Context, customer dto.Customer) error {
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

func (service *GetCustomerByIdUseCase) Execute(ctx context.Context, id uint) (dto.Customer, error) {
	customer, err := service.repository.GetCustomerById(ctx, id)

	if err != nil {
		return dto.Customer{}, responses.GetResponseError(err, "CustomerService")
	}

	return customer, nil
}

func (service *GetCustomerByCPFUseCase) Execute(ctx context.Context, cpf string) (dto.Customer, error) {
	cleanedCPF, validate := service.validateCPFUseCase.Execute(cpf)

	if !validate {
		return dto.Customer{}, &responses.BusinessResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid CPF",
		}
	}

	customer, err := service.repository.GetCustomerByCPF(ctx, cleanedCPF)

	if err != nil {
		return dto.Customer{}, responses.GetResponseError(err, "CustomerService")
	}

	return customer, nil
}

func (uc *LoginCustomerUseCase) Execute(ctx context.Context, cpf string) (dto.Token, error) {
	token, err := uc.repository.Login(ctx, cpf)

	if err != nil {
		return dto.Token{}, responses.GetResponseError(err, "CustomerService")
	}

	return dto.Token{
		AccessToken: token,
	}, nil
}

func (uc *LoginUnknownCustomerUseCase) Execute(ctx context.Context) (dto.Token, error) {
	token, err := uc.repository.LoginUnknown()

	if err != nil {
		return dto.Token{}, responses.GetResponseError(err, "CustomerService")
	}

	return dto.Token{
		AccessToken: token,
	}, nil
}

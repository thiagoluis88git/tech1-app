package handler_test

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
)

type MockCreateCustomerUseCase struct {
	mock.Mock
}

type MockUpdateCustomerUseCase struct {
	mock.Mock
}

type MockGetCustomerByIdUseCase struct {
	mock.Mock
}

type MockGetCustomerByCPFUseCase struct {
	mock.Mock
}

type MockLoginCustomerUseCase struct {
	mock.Mock
}

type MockLoginUnknownCustomerUseCase struct {
	mock.Mock
}

type MockPayOrderUseCase struct {
	mock.Mock
}

type MockGetPaymentTypesUseCase struct {
	mock.Mock
}

func (mock *MockCreateCustomerUseCase) Execute(ctx context.Context, customer dto.Customer) (dto.CustomerResponse, error) {
	args := mock.Called(ctx, customer)
	err := args.Error(1)

	if err != nil {
		return dto.CustomerResponse{}, err
	}

	return args.Get(0).(dto.CustomerResponse), nil
}

func (mock *MockUpdateCustomerUseCase) Execute(ctx context.Context, customer dto.Customer) error {
	args := mock.Called(ctx, customer)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockGetCustomerByIdUseCase) Execute(ctx context.Context, id uint) (dto.Customer, error) {
	args := mock.Called(ctx, id)
	err := args.Error(1)

	if err != nil {
		return dto.Customer{}, err
	}

	return args.Get(0).(dto.Customer), nil
}

func (mock *MockGetCustomerByCPFUseCase) Execute(ctx context.Context, cpf string) (dto.Customer, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return dto.Customer{}, err
	}

	return args.Get(0).(dto.Customer), nil
}

func (mock *MockLoginCustomerUseCase) Execute(ctx context.Context, cpf string) (dto.Token, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return dto.Token{}, err
	}

	return args.Get(0).(dto.Token), nil
}

func (mock *MockLoginUnknownCustomerUseCase) Execute(ctx context.Context) (dto.Token, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return dto.Token{}, err
	}

	return args.Get(0).(dto.Token), nil
}

func (mock *MockPayOrderUseCase) Execute(ctx context.Context, payment dto.Payment) (dto.PaymentResponse, error) {
	args := mock.Called(ctx, payment)
	err := args.Error(1)

	if err != nil {
		return dto.PaymentResponse{}, err
	}

	return args.Get(0).(dto.PaymentResponse), nil
}

func (mock *MockGetPaymentTypesUseCase) Execute() []string {
	args := mock.Called()
	return args.Get(0).([]string)
}

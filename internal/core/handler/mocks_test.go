package handler_test

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
)

type MockCreateCustomerUseCase struct {
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

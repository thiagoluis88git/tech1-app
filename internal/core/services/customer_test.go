package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

var (
	saveCustomer = domain.Customer{
		Name:  "Name",
		CPF:   "12345678900",
		Email: "teste@teste.com",
	}

	customerByCPF = domain.Customer{
		ID:    1,
		Name:  "Name",
		CPF:   "12345678900",
		Email: "teste@teste.com",
	}
)

func TestCustomerServices(t *testing.T) {
	t.Run("got success when creating customer in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockCustomerRepository)
		sut := NewCustomerService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("CreateCustomer", ctx, saveCustomer).Return(uint(1), nil)

		response, err := sut.CreateCustomer(ctx, saveCustomer)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		assert.Equal(t, uint(1), response.Id)
	})

	t.Run("got error when creating customer in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockCustomerRepository)
		sut := NewCustomerService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("CreateCustomer", ctx, saveCustomer).Return(uint(0), &responses.NetworkError{
			Code:    409,
			Message: "Conflict",
		})

		response, err := sut.CreateCustomer(ctx, saveCustomer)

		assert.Error(t, err)
		assert.Empty(t, response)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusConflict, businessError.StatusCode)
	})

	t.Run("got success when updating customer in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockCustomerRepository)
		sut := NewCustomerService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("UpdateCustomer", ctx, saveCustomer).Return(nil)

		err := sut.UpdateCustomer(ctx, saveCustomer)

		assert.NoError(t, err)
	})

	t.Run("got error when updating customer in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockCustomerRepository)
		sut := NewCustomerService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("UpdateCustomer", ctx, saveCustomer).Return(&responses.NetworkError{
			Code:    404,
			Message: "Not Found",
		})

		err := sut.UpdateCustomer(ctx, saveCustomer)

		assert.Error(t, err)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusNotFound, businessError.StatusCode)
	})

	t.Run("got success when getting customer by CPF in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockCustomerRepository)
		sut := NewCustomerService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("GetCustomerByCPF", ctx, "12345678900").Return(customerByCPF, nil)

		response, err := sut.GetCustomerByCPF(ctx, "12345678900")

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		assert.Equal(t, uint(1), response.ID)
		assert.Equal(t, "Name", response.Name)
		assert.Equal(t, "12345678900", response.CPF)
		assert.Equal(t, "teste@teste.com", response.Email)
	})

	t.Run("got error when getting customer by CPF in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockCustomerRepository)
		sut := NewCustomerService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("GetCustomerByCPF", ctx, "12345678900").Return(domain.Customer{}, &responses.NetworkError{
			Code:    404,
			Message: "Not Found",
		})

		response, err := sut.GetCustomerByCPF(ctx, "12345678900")

		assert.Error(t, err)
		assert.Empty(t, response)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusNotFound, businessError.StatusCode)
	})
}

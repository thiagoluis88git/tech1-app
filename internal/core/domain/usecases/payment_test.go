package usecases

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func TestPaymentServices(t *testing.T) {
	t.Run("got success when getting payment types in services", func(t *testing.T) {
		t.Parallel()

		mockPaymentRepo := new(MockPaymentRepository)
		sut := NewGetPaymentTypesUseCasee(mockPaymentRepo)

		mockPaymentRepo.On("GetPaymentTypes").Return([]string{"Crédito", "Débito"})

		response := sut.Execute()

		assert.NotEmpty(t, response)

		assert.Equal(t, 2, len(response))
	})

	t.Run("got success when paying in services", func(t *testing.T) {
		t.Parallel()

		mockPaymentRepo := new(MockPaymentRepository)
		mockPaymentGatewayRepo := new(MockPaymentGatewayRepository)
		sut := NewPayOrderUseCase(mockPaymentRepo, mockPaymentGatewayRepo)

		ctx := context.TODO()

		mockPaymentRepo.On("CreatePaymentOrder", ctx, paymentCreation).Return(paymentResponse, nil)
		mockPaymentGatewayRepo.On("Pay", paymentResponse, paymentCreation).Return(paymentGatewayResponse, nil)
		mockPaymentRepo.On("FinishPaymentWithSuccess", ctx, uint(1)).Return(nil)

		response, err := sut.Execute(ctx, paymentCreation)

		mockPaymentRepo.AssertExpectations(t)
		mockPaymentGatewayRepo.AssertExpectations(t)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("got error when creating payment order paying in services", func(t *testing.T) {
		t.Parallel()

		mockPaymentRepo := new(MockPaymentRepository)
		mockPaymentGatewayRepo := new(MockPaymentGatewayRepository)
		sut := NewPayOrderUseCase(mockPaymentRepo, mockPaymentGatewayRepo)

		ctx := context.TODO()

		mockPaymentRepo.On("CreatePaymentOrder", ctx, paymentCreation).Return(dto.PaymentResponse{}, &responses.LocalError{
			Code:    3,
			Message: "DATABASE_CONFLICT_ERROR",
		})

		response, err := sut.Execute(ctx, paymentCreation)

		mockPaymentRepo.AssertExpectations(t)

		assert.Error(t, err)
		assert.Empty(t, response)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusConflict, businessError.StatusCode)
	})

	t.Run("got error when paying in services", func(t *testing.T) {
		t.Parallel()

		mockPaymentRepo := new(MockPaymentRepository)
		mockPaymentGatewayRepo := new(MockPaymentGatewayRepository)
		sut := NewPayOrderUseCase(mockPaymentRepo, mockPaymentGatewayRepo)

		ctx := context.TODO()

		mockPaymentRepo.On("CreatePaymentOrder", ctx, paymentCreation).Return(paymentResponse, nil)
		mockPaymentGatewayRepo.On("Pay", paymentResponse, paymentCreation).Return(dto.PaymentGatewayResponse{}, &responses.NetworkError{
			Code:    503,
			Message: "Service Unavailable",
		})
		mockPaymentRepo.On("FinishPaymentWithError", ctx, uint(1)).Return(nil)

		response, err := sut.Execute(ctx, paymentCreation)

		mockPaymentRepo.AssertExpectations(t)
		mockPaymentGatewayRepo.AssertExpectations(t)

		assert.Error(t, err)
		assert.Empty(t, response)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusServiceUnavailable, businessError.StatusCode)
	})

	t.Run("got error when finishing payment with success paying in services", func(t *testing.T) {
		t.Parallel()

		mockPaymentRepo := new(MockPaymentRepository)
		mockPaymentGatewayRepo := new(MockPaymentGatewayRepository)
		sut := NewPayOrderUseCase(mockPaymentRepo, mockPaymentGatewayRepo)

		ctx := context.TODO()

		mockPaymentRepo.On("CreatePaymentOrder", ctx, paymentCreation).Return(paymentResponse, nil)
		mockPaymentGatewayRepo.On("Pay", paymentResponse, paymentCreation).Return(paymentGatewayResponse, nil)
		mockPaymentRepo.On("FinishPaymentWithSuccess", ctx, uint(1)).Return(&responses.LocalError{
			Code:    3,
			Message: "DATABASE_CONFLICT_ERROR",
		})

		response, err := sut.Execute(ctx, paymentCreation)

		mockPaymentRepo.AssertExpectations(t)
		mockPaymentGatewayRepo.AssertExpectations(t)

		assert.Error(t, err)
		assert.Empty(t, response)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusConflict, businessError.StatusCode)
	})
}

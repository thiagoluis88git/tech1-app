package usecases

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func newQRCodeOrder() dto.QRCodeOrder {
	return dto.QRCodeOrder{
		TotalPrice: 124.53,
	}
}

func mockQRCodeOrder() dto.Order {
	return dto.Order{
		TotalPrice:   124.53,
		TicketNumber: 3,
		PaymentID:    uint(7),
	}
}

func mockPayment() dto.Payment {
	return dto.Payment{
		TotalPrice:  124.53,
		PaymentType: "QR Code",
	}
}

func mockDate() int64 {
	return time.Now().Unix()
}

func TestQRCodePaymentUseCase(t *testing.T) {
	t.Parallel()

	t.Run("got success when generating QR Code for payment use case", func(t *testing.T) {
		t.Parallel()

		mockOrderRepo := new(MockOrderRepository)
		mockPaymentRepo := new(MockPaymentRepository)
		mockQRCodePaymentepo := new(MockQRCodePaymentRepository)
		sut := NewGenerateQRCodePaymentUseCase(mockQRCodePaymentepo, mockOrderRepo, mockPaymentRepo)

		ctx := context.TODO()
		var wg sync.WaitGroup
		wg.Add(1)
		channel := make(chan bool, 1)
		date := mockDate()

		mockOrderRepo.On("GetNextTicketNumber", ctx, date).Return(3)
		mockPaymentRepo.On("CreatePaymentOrder", ctx, mockPayment()).Return(dto.PaymentResponse{
			PaymentId:        uint(7),
			PaymentGatewayId: "123456",
		}, nil)
		mockOrderRepo.On("CreatePayingOrder", ctx, mockQRCodeOrder()).Return(dto.OrderResponse{
			OrderId: uint(9),
		}, nil)
		mockQRCodePaymentepo.On("Generate", ctx, "token", mockQRCodeOrder(), 9).Return(dto.QRCodeDataResponse{
			Data: "QRCodeData",
		}, nil)

		response, err := sut.Execute(ctx, "token", newQRCodeOrder(), date, &wg, channel)

		mockOrderRepo.AssertExpectations(t)
		mockPaymentRepo.AssertExpectations(t)
		mockQRCodePaymentepo.AssertExpectations(t)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("got error on Generate UseCase when generating QR Code for payment use case", func(t *testing.T) {
		t.Parallel()

		mockOrderRepo := new(MockOrderRepository)
		mockPaymentRepo := new(MockPaymentRepository)
		mockQRCodePaymentepo := new(MockQRCodePaymentRepository)
		sut := NewGenerateQRCodePaymentUseCase(mockQRCodePaymentepo, mockOrderRepo, mockPaymentRepo)

		ctx := context.TODO()
		var wg sync.WaitGroup
		wg.Add(1)
		channel := make(chan bool, 1)
		date := mockDate()

		mockOrderRepo.On("GetNextTicketNumber", ctx, date).Return(3)
		mockPaymentRepo.On("CreatePaymentOrder", ctx, mockPayment()).Return(dto.PaymentResponse{
			PaymentId:        uint(7),
			PaymentGatewayId: "123456",
		}, nil)
		mockOrderRepo.On("CreatePayingOrder", ctx, mockQRCodeOrder()).Return(dto.OrderResponse{
			OrderId: uint(9),
		}, nil)
		mockQRCodePaymentepo.On("Generate", ctx, "token", mockQRCodeOrder(), 9).Return(dto.QRCodeDataResponse{}, &responses.NetworkError{
			Code: 500,
		})
		mockOrderRepo.On("DeleteOrder", ctx, uint(9)).Return(nil)

		response, err := sut.Execute(ctx, "token", newQRCodeOrder(), date, &wg, channel)

		mockOrderRepo.AssertExpectations(t)
		mockPaymentRepo.AssertExpectations(t)
		mockQRCodePaymentepo.AssertExpectations(t)

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got error on Delete Order Repo when generating QR Code for payment use case", func(t *testing.T) {
		t.Parallel()

		mockOrderRepo := new(MockOrderRepository)
		mockPaymentRepo := new(MockPaymentRepository)
		mockQRCodePaymentepo := new(MockQRCodePaymentRepository)
		sut := NewGenerateQRCodePaymentUseCase(mockQRCodePaymentepo, mockOrderRepo, mockPaymentRepo)

		ctx := context.TODO()
		var wg sync.WaitGroup
		wg.Add(1)
		channel := make(chan bool, 1)
		date := mockDate()

		mockOrderRepo.On("GetNextTicketNumber", ctx, date).Return(3)
		mockPaymentRepo.On("CreatePaymentOrder", ctx, mockPayment()).Return(dto.PaymentResponse{
			PaymentId:        uint(7),
			PaymentGatewayId: "123456",
		}, nil)
		mockOrderRepo.On("CreatePayingOrder", ctx, mockQRCodeOrder()).Return(dto.OrderResponse{
			OrderId: uint(9),
		}, nil)
		mockQRCodePaymentepo.On("Generate", ctx, "token", mockQRCodeOrder(), 9).Return(dto.QRCodeDataResponse{}, &responses.NetworkError{
			Code: 500,
		})
		mockOrderRepo.On("DeleteOrder", ctx, uint(9)).Return(&responses.LocalError{
			Code: responses.DATABASE_CONSTRAINT_ERROR,
		})

		response, err := sut.Execute(ctx, "token", newQRCodeOrder(), date, &wg, channel)

		mockOrderRepo.AssertExpectations(t)
		mockPaymentRepo.AssertExpectations(t)
		mockQRCodePaymentepo.AssertExpectations(t)

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got error on CreatePayingOrder Repo when generating QR Code for payment use case", func(t *testing.T) {
		t.Parallel()

		mockOrderRepo := new(MockOrderRepository)
		mockPaymentRepo := new(MockPaymentRepository)
		mockQRCodePaymentepo := new(MockQRCodePaymentRepository)
		sut := NewGenerateQRCodePaymentUseCase(mockQRCodePaymentepo, mockOrderRepo, mockPaymentRepo)

		ctx := context.TODO()
		var wg sync.WaitGroup
		wg.Add(1)
		channel := make(chan bool, 1)
		date := mockDate()

		mockOrderRepo.On("GetNextTicketNumber", ctx, date).Return(3)
		mockPaymentRepo.On("CreatePaymentOrder", ctx, mockPayment()).Return(dto.PaymentResponse{
			PaymentId:        uint(7),
			PaymentGatewayId: "123456",
		}, nil)
		mockOrderRepo.On("CreatePayingOrder", ctx, mockQRCodeOrder()).Return(dto.OrderResponse{}, &responses.LocalError{
			Code: responses.DATABASE_CONSTRAINT_ERROR,
		})

		response, err := sut.Execute(ctx, "token", newQRCodeOrder(), date, &wg, channel)

		mockOrderRepo.AssertExpectations(t)
		mockPaymentRepo.AssertExpectations(t)
		mockQRCodePaymentepo.AssertExpectations(t)

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got error on CreatePayingOrder Repo when generating QR Code for payment use case", func(t *testing.T) {
		t.Parallel()

		mockOrderRepo := new(MockOrderRepository)
		mockPaymentRepo := new(MockPaymentRepository)
		mockQRCodePaymentepo := new(MockQRCodePaymentRepository)
		sut := NewGenerateQRCodePaymentUseCase(mockQRCodePaymentepo, mockOrderRepo, mockPaymentRepo)

		ctx := context.TODO()
		var wg sync.WaitGroup
		wg.Add(1)
		channel := make(chan bool, 1)
		date := mockDate()

		mockOrderRepo.On("GetNextTicketNumber", ctx, date).Return(3)
		mockPaymentRepo.On("CreatePaymentOrder", ctx, mockPayment()).Return(dto.PaymentResponse{}, &responses.NetworkError{
			Code: 500,
		})

		response, err := sut.Execute(ctx, "token", newQRCodeOrder(), date, &wg, channel)

		mockOrderRepo.AssertExpectations(t)
		mockPaymentRepo.AssertExpectations(t)
		mockQRCodePaymentepo.AssertExpectations(t)

		assert.Error(t, err)
		assert.Empty(t, response)
	})
}

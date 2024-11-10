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

func newExternalPaymentEvent() dto.ExternalPaymentEvent {
	return dto.ExternalPaymentEvent{
		Resource: "Resource",
		Topic:    "merchant_order",
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

	t.Run("got success when finishing order payment use case", func(t *testing.T) {
		t.Parallel()

		mockOrderRepo := new(MockOrderRepository)
		mockPaymentRepo := new(MockPaymentRepository)
		mockQRCodePaymentepo := new(MockQRCodePaymentRepository)
		sut := NewFinishOrderForQRCodeUseCase(mockQRCodePaymentepo, mockOrderRepo, mockPaymentRepo)

		ctx := context.TODO()

		mockQRCodePaymentepo.On("GetQRCodePaymentData", ctx, "token", "Resource").Return(dto.ExternalPaymentInformation{
			ID:                int64(2),
			Status:            "COMPLETE",
			OrderStatus:       "paid",
			ExternalReference: "123|789",
		}, nil)

		mockPaymentRepo.On("FinishPaymentWithSuccess", ctx, uint(789)).Return(nil)
		mockOrderRepo.On("FinishOrderWithPayment", ctx, uint(123), uint(789)).Return(nil)
		err := sut.Execute(ctx, "token", newExternalPaymentEvent())

		mockOrderRepo.AssertExpectations(t)
		mockPaymentRepo.AssertExpectations(t)
		mockQRCodePaymentepo.AssertExpectations(t)

		assert.NoError(t, err)
	})

	t.Run("got no error on not paid when finishing order payment use case", func(t *testing.T) {
		t.Parallel()

		mockOrderRepo := new(MockOrderRepository)
		mockPaymentRepo := new(MockPaymentRepository)
		mockQRCodePaymentepo := new(MockQRCodePaymentRepository)
		sut := NewFinishOrderForQRCodeUseCase(mockQRCodePaymentepo, mockOrderRepo, mockPaymentRepo)

		ctx := context.TODO()

		mockQRCodePaymentepo.On("GetQRCodePaymentData", ctx, "token", "Resource").Return(dto.ExternalPaymentInformation{
			ID:                int64(2),
			Status:            "COMPLETE",
			OrderStatus:       "waiting",
			ExternalReference: "123|789",
		}, nil)

		err := sut.Execute(ctx, "token", newExternalPaymentEvent())

		mockQRCodePaymentepo.AssertExpectations(t)

		assert.NoError(t, err)
	})

	t.Run("got error on GetQRCodePaymentData Repository when finishing order payment use case", func(t *testing.T) {
		t.Parallel()

		mockOrderRepo := new(MockOrderRepository)
		mockPaymentRepo := new(MockPaymentRepository)
		mockQRCodePaymentepo := new(MockQRCodePaymentRepository)
		sut := NewFinishOrderForQRCodeUseCase(mockQRCodePaymentepo, mockOrderRepo, mockPaymentRepo)

		ctx := context.TODO()

		mockQRCodePaymentepo.On("GetQRCodePaymentData", ctx, "token", "Resource").
			Return(dto.ExternalPaymentInformation{}, &responses.NetworkError{
				Code: 500,
			})

		err := sut.Execute(ctx, "token", newExternalPaymentEvent())

		mockQRCodePaymentepo.AssertExpectations(t)

		assert.Error(t, err)
	})
}

package services

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func TestOrderServices(t *testing.T) {
	t.Run("got success when generating ticket number in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		date := time.Now().UnixMilli()

		mockRepo.On("GetNextTicketNumber", ctx, date).Return(1, nil)

		response := sut.GenerateTicket(ctx, date)

		assert.Equal(t, 1, response)
	})

	t.Run("got success when creating order in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		date := time.Now().UnixMilli()

		mockRepo.On("CreateOrder", ctx, orderCreation).Return(orderCreationResponse, nil)
		mockRepo.On("GetNextTicketNumber", ctx, date).Return(1, nil)

		wg := &sync.WaitGroup{}
		ch := make(chan bool, 1)

		wg.Add(1)
		response, err := sut.CreateOrder(ctx, orderCreation, date, wg, ch)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("got success when creating order with customer in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		date := time.Now().UnixMilli()

		mockRepo.On("CreateOrder", ctx, orderCreationWithCustomer).Return(orderWithCustomerCreationResponse, nil)
		mockRepo.On("GetNextTicketNumber", ctx, date).Return(1, nil)
		mockCustomerRepo.On("GetCustomerById", ctx, customerId).Return(customerResponse, nil)

		wg := &sync.WaitGroup{}
		ch := make(chan bool, 1)

		wg.Add(1)
		response, err := sut.CreateOrder(ctx, orderCreationWithCustomer, date, wg, ch)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("got error when creating order in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		date := time.Now().UnixMilli()

		mockRepo.On("GetNextTicketNumber", ctx, date).Return(1, nil)
		mockRepo.On("CreateOrder", ctx, orderCreationWithCustomer).Return(domain.OrderResponse{}, &responses.NetworkError{
			Code:    409,
			Message: "Conflict",
		})

		wg := &sync.WaitGroup{}
		ch := make(chan bool, 1)

		wg.Add(1)
		response, err := sut.CreateOrder(ctx, orderCreationWithCustomer, date, wg, ch)

		assert.Error(t, err)
		assert.Empty(t, response)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusConflict, businessError.StatusCode)
	})

	t.Run("got success when getting order by id in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrderById", ctx, uint(1)).Return(orderCreationResponse, nil)

		response, err := sut.GetOrderById(ctx, uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("got error when getting order by id in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrderById", ctx, uint(1)).Return(domain.OrderProductResponse{}, &responses.NetworkError{
			Code:    404,
			Message: "Not Found",
		})

		response, err := sut.GetOrderById(ctx, uint(1))

		assert.Error(t, err)
		assert.Empty(t, response)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusNotFound, businessError.StatusCode)
	})

	t.Run("got success when getting orders to prepare in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrderById", ctx, uint(1)).Return(domain.OrderResponse{
			OrderStatus: "Criado",
		}, nil)
		mockRepo.On("GetOrdersToPrepare", ctx).Return(ordersList, nil)

		response, err := sut.GetOrdersToPrepare(ctx)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		assert.Equal(t, 1, len(response))
	})

	t.Run("got error when getting orders to prepare in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrderById", ctx, uint(1)).Return(domain.OrderResponse{
			OrderStatus: "Criado",
		}, nil)
		mockRepo.On("GetOrdersToPrepare", ctx).Return(domain.OrderResponse{}, &responses.NetworkError{
			Code:    404,
			Message: "Not Found",
		})

		response, err := sut.GetOrdersToPrepare(ctx)

		assert.Error(t, err)
		assert.Empty(t, response)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusNotFound, businessError.StatusCode)
	})

	t.Run("got success when getting orders status in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrdersToFollow", ctx).Return(ordersList, nil)

		response, err := sut.GetOrdersToFollow(ctx)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		assert.Equal(t, 1, len(response))
	})

	t.Run("got error when getting orders status in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrdersToFollow", ctx).Return(domain.OrderResponse{}, &responses.NetworkError{
			Code:    404,
			Message: "Not Found",
		})

		response, err := sut.GetOrdersToFollow(ctx)

		assert.Error(t, err)
		assert.Empty(t, response)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusNotFound, businessError.StatusCode)
	})

	t.Run("got success when updating order to delivered in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrderById", ctx, uint(1)).Return(domain.OrderResponse{
			OrderStatus: "Finalizado",
		}, nil)
		mockRepo.On("UpdateToDelivered", ctx, uint(1)).Return(nil)

		err := sut.UpdateToDelivered(ctx, uint(1))

		assert.NoError(t, err)
	})

	t.Run("got error when updating order to delivered in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrderById", ctx, uint(1)).Return(domain.OrderResponse{
			OrderStatus: "Finalizado",
		}, nil)
		mockRepo.On("UpdateToDelivered", ctx, uint(1)).Return(&responses.NetworkError{
			Code:    404,
			Message: "Not Found",
		})

		err := sut.UpdateToDelivered(ctx, uint(1))

		assert.Error(t, err)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusNotFound, businessError.StatusCode)
	})

	t.Run("got success when updating order to done in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrderById", ctx, uint(1)).Return(domain.OrderResponse{
			OrderStatus: "Preparando",
		}, nil)
		mockRepo.On("UpdateToDone", ctx, uint(1)).Return(nil)

		err := sut.UpdateToDone(ctx, uint(1))

		assert.NoError(t, err)
	})

	t.Run("got error when updating order to done in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrderById", ctx, uint(1)).Return(domain.OrderResponse{
			OrderStatus: "Preparando",
		}, nil)
		mockRepo.On("UpdateToDone", ctx, uint(1)).Return(&responses.NetworkError{
			Code:    404,
			Message: "Not Found",
		})

		err := sut.UpdateToDone(ctx, uint(1))

		assert.Error(t, err)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusNotFound, businessError.StatusCode)
	})

	t.Run("got success when updating order to not delivered in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrderById", ctx, uint(1)).Return(domain.OrderResponse{
			OrderStatus: "Finalizado",
		}, nil)
		mockRepo.On("UpdateToNotDelivered", ctx, uint(1)).Return(nil)

		err := sut.UpdateToNotDelivered(ctx, uint(1))

		assert.NoError(t, err)
	})

	t.Run("got error when updating order to not delivered in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrderById", ctx, uint(1)).Return(domain.OrderResponse{
			OrderStatus: "Finalizado",
		}, nil)
		mockRepo.On("UpdateToNotDelivered", ctx, uint(1)).Return(&responses.NetworkError{
			Code:    404,
			Message: "Not Found",
		})

		err := sut.UpdateToNotDelivered(ctx, uint(1))

		assert.Error(t, err)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusNotFound, businessError.StatusCode)
	})

	t.Run("got success when updating order to preparing in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("GetOrderById", ctx, uint(1)).Return(domain.OrderResponse{
			OrderStatus: "Criado",
		}, nil)
		mockRepo.On("UpdateToPreparing", ctx, uint(1)).Return(nil)

		err := sut.UpdateToPreparing(ctx, uint(1))

		assert.NoError(t, err)
	})

	t.Run("got error when updating order to preparing in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockOrderRepository)
		mockCustomerRepo := new(MockCustomerRepository)
		validateToPrepare := NewValidateOrderToPrepareUseCase(mockRepo)
		validateToDone := NewValidateOrderToDoneUseCase(mockRepo)
		validateToDeliveredOrNot := NewValidateOrderToDeliveredOrNotUseCase(mockRepo)

		sut := NewOrderService(mockRepo, mockCustomerRepo, validateToPrepare, validateToDone, validateToDeliveredOrNot)

		ctx := context.TODO()

		mockRepo.On("UpdateToPreparing", ctx, uint(1)).Return(&responses.NetworkError{
			Code:    404,
			Message: "Not Found",
		})

		mockRepo.On("GetOrderById", ctx, uint(1)).Return(domain.OrderResponse{
			OrderStatus: "Criado",
		}, nil)
		err := sut.UpdateToPreparing(ctx, uint(1))

		assert.Error(t, err)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusNotFound, businessError.StatusCode)
	})
}

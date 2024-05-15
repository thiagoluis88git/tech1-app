package services

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

var (
	orderCreation = domain.Order{
		TotalPrice:   12345,
		PaymentID:    1,
		TicketNumber: 1,
		OrderProduct: []domain.OrderProduct{
			{
				ProductID: 1,
			},
			{
				ProductID: 2,
			},
		},
	}
	customerId = uint(1)

	orderCreationWithCustomer = domain.Order{
		TotalPrice:   12345,
		PaymentID:    1,
		TicketNumber: 1,
		CustomerID:   &customerId,
		OrderProduct: []domain.OrderProduct{
			{
				ProductID: 1,
			},
			{
				ProductID: 2,
			},
		},
	}

	orderCreationResponse = domain.OrderResponse{
		OrderId:      1,
		OrderDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		TicketNumber: 1,
		OrderStatus:  "Criado",
		OrderProduct: []domain.OrderProductResponse{
			{
				ProductID:   1,
				ProductName: "ProductName 1",
			},
			{
				ProductID:   2,
				ProductName: "ProductName 2",
			},
		},
	}

	customerName                      = "Customer Name"
	orderWithCustomerCreationResponse = domain.OrderResponse{
		OrderId:      1,
		OrderDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
		TicketNumber: 1,
		OrderStatus:  "Criado",
		CustomerName: &customerName,
		OrderProduct: []domain.OrderProductResponse{
			{
				ProductID:   1,
				ProductName: "ProductName 1",
			},
			{
				ProductID:   2,
				ProductName: "ProductName 2",
			},
		},
	}

	customerResponse = domain.Customer{
		Name: "Customer",
	}

	ordersList = []domain.OrderResponse{
		{
			OrderId:      1,
			OrderDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			TicketNumber: 1,
			OrderStatus:  "Criado",
			OrderProduct: []domain.OrderProductResponse{
				{
					ProductID:   1,
					ProductName: "ProductName 1",
				},
				{
					ProductID:   2,
					ProductName: "ProductName 2",
				},
			},
		},
	}

	paymentCreation = domain.Payment{
		TotalPrice:  1234,
		PaymentType: "Crédito",
	}

	paymentResponse = domain.PaymentResponse{
		PaymentId:        1,
		PaymentGatewayId: "123",
		PaymentDate:      time.Date(2024, 10, 10, 0, 0, 0, 0, time.Local),
	}

	paymentGatewayResponse = domain.PaymentGatewayResponse{
		PaymentGatewayId: "1234",
		PaymentDate:      time.Date(2024, 10, 10, 0, 0, 0, 0, time.Local),
	}
)

type MockOrderRepository struct {
	mock.Mock
}

type MockCustomerRepository struct {
	mock.Mock
}

type MockPaymentRepository struct {
	mock.Mock
}

type MockPaymentGatewayRepository struct {
	mock.Mock
}

func (mock *MockCustomerRepository) CreateCustomer(ctx context.Context, customer domain.Customer) (uint, error) {
	args := mock.Called(ctx, customer)
	err := args.Error(1)

	if err != nil {
		return 0, err
	}

	return args.Get(0).(uint), nil
}

func (mock *MockCustomerRepository) UpdateCustomer(ctx context.Context, customer domain.Customer) error {
	args := mock.Called(ctx, customer)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockCustomerRepository) GetCustomerByCPF(ctx context.Context, cpf string) (domain.Customer, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return domain.Customer{}, err
	}

	return args.Get(0).(domain.Customer), nil
}

func (mock *MockCustomerRepository) GetCustomerById(ctx context.Context, id uint) (domain.Customer, error) {
	args := mock.Called(ctx, id)
	err := args.Error(1)

	if err != nil {
		return domain.Customer{}, err
	}

	return args.Get(0).(domain.Customer), nil
}

func (mock *MockOrderRepository) CreateOrder(ctx context.Context, order domain.Order) (domain.OrderResponse, error) {
	args := mock.Called(ctx, order)
	err := args.Error(1)

	if err != nil {
		return domain.OrderResponse{}, err
	}

	return args.Get(0).(domain.OrderResponse), nil
}

func (mock *MockOrderRepository) GetOrderById(ctx context.Context, orderId uint) (domain.OrderResponse, error) {
	args := mock.Called(ctx, orderId)
	err := args.Error(1)

	if err != nil {
		return domain.OrderResponse{}, err
	}

	return args.Get(0).(domain.OrderResponse), nil
}

func (mock *MockOrderRepository) GetOrdersToPrepare(ctx context.Context) ([]domain.OrderResponse, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return []domain.OrderResponse{}, err
	}

	return args.Get(0).([]domain.OrderResponse), nil
}

func (mock *MockOrderRepository) GetOrdersStatus(ctx context.Context) ([]domain.OrderResponse, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return []domain.OrderResponse{}, err
	}

	return args.Get(0).([]domain.OrderResponse), nil
}

func (mock *MockOrderRepository) UpdateToPreparing(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) UpdateToDone(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) UpdateToDelivered(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) UpdateToNotDelivered(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) GetNextTicketNumber(ctx context.Context, date int64) int {
	args := mock.Called(ctx, date)
	return args.Get(0).(int)
}

func (mock *MockPaymentRepository) GetPaymentTypes() []string {
	args := mock.Called()
	return args.Get(0).([]string)
}

func (mock *MockPaymentRepository) CreatePaymentOrder(ctx context.Context, payment domain.Payment) (domain.PaymentResponse, error) {
	args := mock.Called(ctx, payment)
	err := args.Error(1)

	if err != nil {
		return domain.PaymentResponse{}, err
	}

	return args.Get(0).(domain.PaymentResponse), nil
}

func (mock *MockPaymentRepository) FinishPaymentWithSuccess(ctx context.Context, paymentId uint) error {
	args := mock.Called(ctx, paymentId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockPaymentRepository) FinishPaymentWithError(ctx context.Context, paymentId uint) error {
	args := mock.Called(ctx, paymentId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockPaymentGatewayRepository) Pay(paymentResonse domain.PaymentResponse, payment domain.Payment) (domain.PaymentGatewayResponse, error) {
	args := mock.Called(paymentResonse, payment)
	err := args.Error(1)

	if err != nil {
		return domain.PaymentGatewayResponse{}, err
	}

	return args.Get(0).(domain.PaymentGatewayResponse), nil
}

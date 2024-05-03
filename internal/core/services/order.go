package services

import (
	"context"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/ports"
	"thiagoluis88git/tech1/pkg/responses"
)

type OrderService struct {
	orderRepo    ports.OrderRepository
	customerRepo ports.CustomerRepository
}

func NewOrderService(
	orderRepo ports.OrderRepository,
	customerRepo ports.CustomerRepository,
) *OrderService {
	return &OrderService{
		orderRepo:    orderRepo,
		customerRepo: customerRepo,
	}
}

func (service *OrderService) CreateOrder(ctx context.Context, order domain.Order) (domain.OrderResponse, error) {
	response, err := service.orderRepo.CreateOrder(ctx, order)

	if err != nil {
		return domain.OrderResponse{}, responses.GetResponseError(err, "OrderService")
	}

	if order.CustomerID != nil {
		customer, err := service.customerRepo.GetCustomerById(ctx, *order.CustomerID)
		if err == nil {
			response.CustomerName = &customer.Name
		}
	}

	return response, nil
}

func (service *OrderService) GetOrderById(ctx context.Context, orderId uint) (domain.OrderResponse, error) {
	response, err := service.orderRepo.GetOrderById(ctx, orderId)

	if err != nil {
		return domain.OrderResponse{}, responses.GetResponseError(err, "OrderService")
	}

	return response, nil
}

func (service *OrderService) GetOrdersToPrepare(ctx context.Context) ([]domain.OrderResponse, error) {
	response, err := service.orderRepo.GetOrdersToPrepare(ctx)

	if err != nil {
		return []domain.OrderResponse{}, responses.GetResponseError(err, "OrderService")
	}

	return response, nil
}

func (service *OrderService) UpdateToPreparing(ctx context.Context, orderId uint) error {
	err := service.orderRepo.UpdateToPreparing(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService")
	}

	return nil
}

func (service *OrderService) UpdateToDone(ctx context.Context, orderId uint) error {
	err := service.orderRepo.UpdateToDone(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService")
	}

	return nil
}

func (service *OrderService) UpdateToDelivered(ctx context.Context, orderId uint) error {
	err := service.orderRepo.UpdateToDelivered(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService")
	}

	return nil
}

func (service *OrderService) UpdateToNotDelivered(ctx context.Context, orderId uint) error {
	err := service.orderRepo.UpdateToNotDelivered(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "OrderService")
	}

	return nil
}

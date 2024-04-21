package services

import (
	"context"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/ports"
	"thiagoluis88git/tech1/pkg/responses"
)

type OrderService struct {
	repository ports.OrderRepository
}

func NewOrderService(repository ports.OrderRepository) *OrderService {
	return &OrderService{
		repository: repository,
	}
}

func (service *OrderService) CreateOrder(ctx context.Context, order domain.Order) (domain.OrderResponse, error) {
	response, err := service.repository.CreateOrder(ctx, order)

	if err != nil {
		return domain.OrderResponse{}, responses.GetResponseError(err, "OrderService")
	}

	return response, nil
}

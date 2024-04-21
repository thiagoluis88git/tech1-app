package ports

import (
	"context"
	"thiagoluis88git/tech1/internal/core/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order domain.Order) (domain.OrderResponse, error)
	FinishOrderPayment(ctx context.Context, orderId uint) error
}

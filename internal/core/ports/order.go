package ports

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order domain.Order) (domain.OrderResponse, error)
	GetOrderById(ctx context.Context, orderId uint) (domain.OrderResponse, error)
	GetOrdersToPrepare(ctx context.Context) ([]domain.OrderResponse, error)
	GetOrdersStatus(ctx context.Context) ([]domain.OrderResponse, error)
	UpdateToPreparing(ctx context.Context, orderId uint) error
	UpdateToDone(ctx context.Context, orderId uint) error
	UpdateToDelivered(ctx context.Context, orderId uint) error
	UpdateToNotDelivered(ctx context.Context, orderId uint) error
	GetNextTicketNumber(ctx context.Context, date int64) int
}

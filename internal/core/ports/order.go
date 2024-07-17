package ports

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order domain.Order) (domain.OrderResponse, error)
	CreatePayingOrder(ctx context.Context, order domain.Order) (domain.OrderResponse, error)
	FinishOrderWithPayment(ctx context.Context, orderID uint, paymentID uint) error
	DeleteOrder(ctx context.Context, orderID uint) error
	GetOrderById(ctx context.Context, orderID uint) (domain.OrderResponse, error)
	GetOrdersToPrepare(ctx context.Context) ([]domain.OrderResponse, error)
	GetOrdersToFollow(ctx context.Context) ([]domain.OrderResponse, error)
	GetOrdersWaitingPayment(ctx context.Context) ([]domain.OrderResponse, error)
	UpdateToPreparing(ctx context.Context, orderID uint) error
	UpdateToDone(ctx context.Context, orderID uint) error
	UpdateToDelivered(ctx context.Context, orderID uint) error
	UpdateToNotDelivered(ctx context.Context, orderID uint) error
	GetNextTicketNumber(ctx context.Context, date int64) int
}

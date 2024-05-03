package ports

import (
	"context"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

type PaymentRepository interface {
	GetPaymentTypes() []string
	CreatePaymentOrder(ctx context.Context, payment domain.Payment) (domain.PaymentResponse, error)
	FinishPaymentWithSuccess(ctx context.Context, paymentId uint) error
	FinishPaymentWithError(ctx context.Context, paymentId uint) error
}

package repository

import (
	"context"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
)

type PaymentRepository interface {
	GetPaymentTypes() []string
	CreatePaymentOrder(ctx context.Context, payment dto.Payment) (dto.PaymentResponse, error)
	FinishPaymentWithSuccess(ctx context.Context, paymentId uint) error
	FinishPaymentWithError(ctx context.Context, paymentId uint) error
}

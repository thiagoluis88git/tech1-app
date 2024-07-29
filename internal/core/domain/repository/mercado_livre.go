package repository

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
)

type QRCodePaymentRepository interface {
	Generate(ctx context.Context, token string, form dto.Order, orderID int) (dto.QRCodeDataResponse, error)
	GetQRCodePaymentData(ctx context.Context, token string, endpoint string) (dto.ExternalPaymentInformation, error)
}

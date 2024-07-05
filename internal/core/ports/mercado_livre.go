package ports

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

type QRCodePaymentRepository interface {
	Generate(ctx context.Context, token string, form domain.Order, orderID int) (domain.QRCodeDataResponse, error)
	GetQRCodePaymentData(ctx context.Context, token string, endpoint string) (domain.MercadoLivrePayment, error)
}

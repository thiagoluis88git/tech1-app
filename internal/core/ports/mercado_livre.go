package ports

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

type MercadoLivreRepository interface {
	Generate(ctx context.Context, token string, form domain.Order, orderID int) (domain.QRCodeDataResponse, error)
	GetMercadoLivrePaymentData(ctx context.Context, token string, endpoint string) error
}

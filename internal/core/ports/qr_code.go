package ports

import "github.com/thiagoluis88git/tech1/internal/core/domain"

type QRCodeGeneratorRepository interface {
	Generate(token string, form domain.Order, orderID int) (domain.QRCodeDataResponse, error)
}

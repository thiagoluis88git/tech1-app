package ports

import "github.com/thiagoluis88git/tech1/internal/core/domain"

type QRCodeGeneratorRepository interface {
	Generate(token string, form domain.QRCodeForm) (domain.QRCodeDataResponse, error)
}

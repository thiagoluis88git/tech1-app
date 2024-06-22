package services

import (
	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type QRCodeGeneratorService struct {
	repository ports.QRCodeGeneratorRepository
}

func NewQRCodeGeneratorService(repository ports.QRCodeGeneratorRepository) *QRCodeGeneratorService {
	return &QRCodeGeneratorService{
		repository: repository,
	}
}

func (service *QRCodeGeneratorService) GenerateQRCode(token string, form domain.QRCodeForm) (domain.QRCodeDataResponse, error) {
	qrCode, err := service.repository.Generate(token, form)

	if err != nil {
		return domain.QRCodeDataResponse{}, responses.GetResponseError(err, "QRCodeGeneratorService")
	}

	return qrCode, nil
}

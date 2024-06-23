package services

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type QRCodeGeneratorService struct {
	repository      ports.QRCodeGeneratorRepository
	orderRepository ports.OrderRepository
}

func NewQRCodeGeneratorService(
	repository ports.QRCodeGeneratorRepository,
	orderRepository ports.OrderRepository,
) *QRCodeGeneratorService {
	return &QRCodeGeneratorService{
		repository:      repository,
		orderRepository: orderRepository,
	}
}

func (service *QRCodeGeneratorService) GenerateQRCode(
	ctx context.Context,
	token string,
	order domain.Order,
) (domain.QRCodeDataResponse, error) {
	orderResponse, err := service.orderRepository.CreatePayingOrder(ctx, order)

	if err != nil {
		return domain.QRCodeDataResponse{}, responses.GetResponseError(err, "QRCodeGeneratorService")
	}

	qrCode, err := service.repository.Generate(token, order, int(orderResponse.OrderId))

	if err != nil {
		errDelete := service.orderRepository.DeleteOrder(ctx, orderResponse.OrderId)

		if errDelete != nil {
			return domain.QRCodeDataResponse{}, responses.GetResponseError(errDelete, "QRCodeGeneratorService")
		}

		return domain.QRCodeDataResponse{}, responses.GetResponseError(err, "QRCodeGeneratorService")
	}

	return qrCode, nil
}

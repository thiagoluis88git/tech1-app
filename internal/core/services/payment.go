package services

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type PaymentService struct {
	paymentRepo    ports.PaymentRepository
	paymentGateway ports.PaymentGateway
}

func NewPaymentService(paymentRepo ports.PaymentRepository, paymentGateway ports.PaymentGateway) *PaymentService {
	return &PaymentService{
		paymentRepo:    paymentRepo,
		paymentGateway: paymentGateway,
	}
}

func (service *PaymentService) PayOrder(ctx context.Context, payment domain.Payment) (domain.PaymentResponse, error) {
	paymentResponse, err := service.paymentRepo.CreatePaymentOrder(ctx, payment)

	if err != nil {
		return domain.PaymentResponse{}, responses.GetResponseError(err, "PaymentService")
	}

	gatewayResponse, err := service.paymentGateway.Pay(paymentResponse, payment)

	if err != nil {
		paymentWithError := service.paymentRepo.FinishPaymentWithError(ctx, paymentResponse.PaymentId)

		if paymentWithError != nil {
			return domain.PaymentResponse{}, responses.GetResponseError(paymentWithError, "PaymentService")
		}

		return domain.PaymentResponse{}, responses.GetResponseError(err, "PaymentService")
	}

	err = service.paymentRepo.FinishPaymentWithSuccess(ctx, paymentResponse.PaymentId)

	if err != nil {
		return domain.PaymentResponse{}, responses.GetResponseError(err, "PaymentService")
	}

	return domain.PaymentResponse{
		PaymentId:        paymentResponse.PaymentId,
		PaymentGatewayId: gatewayResponse.PaymentGatewayId,
		PaymentDate:      gatewayResponse.PaymentDate,
	}, nil
}

func (service *PaymentService) GetPaymentTypes() []string {
	return service.paymentRepo.GetPaymentTypes()
}

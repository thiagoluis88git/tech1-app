package usecases

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type PayOrderUseCase struct {
	paymentRepo    ports.PaymentRepository
	paymentGateway ports.PaymentGateway
}

type GetPaymentTypesUseCase struct {
	paymentRepo ports.PaymentRepository
}

func NewPayOrderUseCase(paymentRepo ports.PaymentRepository, paymentGateway ports.PaymentGateway) *PayOrderUseCase {
	return &PayOrderUseCase{
		paymentRepo:    paymentRepo,
		paymentGateway: paymentGateway,
	}
}

func NewGetPaymentTypesUseCasee(paymentRepo ports.PaymentRepository) *GetPaymentTypesUseCase {
	return &GetPaymentTypesUseCase{
		paymentRepo: paymentRepo,
	}
}

func (usecase *PayOrderUseCase) Execute(ctx context.Context, payment domain.Payment) (domain.PaymentResponse, error) {
	paymentResponse, err := usecase.paymentRepo.CreatePaymentOrder(ctx, payment)

	if err != nil {
		return domain.PaymentResponse{}, responses.GetResponseError(err, "PaymentService")
	}

	gatewayResponse, err := usecase.paymentGateway.Pay(paymentResponse, payment)

	if err != nil {
		paymentWithError := usecase.paymentRepo.FinishPaymentWithError(ctx, paymentResponse.PaymentId)

		if paymentWithError != nil {
			return domain.PaymentResponse{}, responses.GetResponseError(paymentWithError, "PaymentService")
		}

		return domain.PaymentResponse{}, responses.GetResponseError(err, "PaymentService")
	}

	err = usecase.paymentRepo.FinishPaymentWithSuccess(ctx, paymentResponse.PaymentId)

	if err != nil {
		return domain.PaymentResponse{}, responses.GetResponseError(err, "PaymentService")
	}

	return domain.PaymentResponse{
		PaymentId:        paymentResponse.PaymentId,
		PaymentGatewayId: gatewayResponse.PaymentGatewayId,
		PaymentDate:      gatewayResponse.PaymentDate,
	}, nil
}

func (usecase *GetPaymentTypesUseCase) Execute() []string {
	return usecase.paymentRepo.GetPaymentTypes()
}

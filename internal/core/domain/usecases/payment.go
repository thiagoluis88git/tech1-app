package usecases

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/domain/repository"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type PayOrderUseCase interface {
	Execute(ctx context.Context, payment dto.Payment) (dto.PaymentResponse, error)
}

type PayOrderUseCaseImpl struct {
	paymentRepo    repository.PaymentRepository
	paymentGateway repository.PaymentGateway
}

type GetPaymentTypesUseCase interface {
	Execute() []string
}

type GetPaymentTypesUseCaseImpl struct {
	paymentRepo repository.PaymentRepository
}

func NewPayOrderUseCase(paymentRepo repository.PaymentRepository, paymentGateway repository.PaymentGateway) PayOrderUseCase {
	return &PayOrderUseCaseImpl{
		paymentRepo:    paymentRepo,
		paymentGateway: paymentGateway,
	}
}

func NewGetPaymentTypesUseCasee(paymentRepo repository.PaymentRepository) GetPaymentTypesUseCase {
	return &GetPaymentTypesUseCaseImpl{
		paymentRepo: paymentRepo,
	}
}

func (usecase *PayOrderUseCaseImpl) Execute(ctx context.Context, payment dto.Payment) (dto.PaymentResponse, error) {
	paymentResponse, err := usecase.paymentRepo.CreatePaymentOrder(ctx, payment)

	if err != nil {
		return dto.PaymentResponse{}, responses.GetResponseError(err, "PaymentService")
	}

	gatewayResponse, err := usecase.paymentGateway.Pay(paymentResponse, payment)

	if err != nil {
		paymentWithError := usecase.paymentRepo.FinishPaymentWithError(ctx, paymentResponse.PaymentId)

		if paymentWithError != nil {
			return dto.PaymentResponse{}, responses.GetResponseError(paymentWithError, "PaymentService")
		}

		return dto.PaymentResponse{}, responses.GetResponseError(err, "PaymentService")
	}

	err = usecase.paymentRepo.FinishPaymentWithSuccess(ctx, paymentResponse.PaymentId)

	if err != nil {
		return dto.PaymentResponse{}, responses.GetResponseError(err, "PaymentService")
	}

	return dto.PaymentResponse{
		PaymentId:        paymentResponse.PaymentId,
		PaymentGatewayId: gatewayResponse.PaymentGatewayId,
		PaymentDate:      gatewayResponse.PaymentDate,
	}, nil
}

func (usecase *GetPaymentTypesUseCaseImpl) Execute() []string {
	return usecase.paymentRepo.GetPaymentTypes()
}

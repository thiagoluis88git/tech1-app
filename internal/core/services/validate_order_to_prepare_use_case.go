package services

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type ValidateOrderToPrepareUseCase struct {
	repository ports.OrderRepository
}

func NewValidateOrderToPrepareUseCase(repository ports.OrderRepository) *ValidateOrderToPrepareUseCase {
	return &ValidateOrderToPrepareUseCase{
		repository: repository,
	}
}

func (usecase *ValidateOrderToPrepareUseCase) Execute(ctx context.Context, orderId uint) error {
	response, err := usecase.repository.GetOrderById(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "ValidateOrderToPrepareUseCase -> GetOrderById")
	}

	if response.OrderStatus != "Criado" {
		return &responses.BusinessResponse{
			StatusCode: 428,
			Message:    "The order must be in Criado status",
		}
	}

	return nil
}

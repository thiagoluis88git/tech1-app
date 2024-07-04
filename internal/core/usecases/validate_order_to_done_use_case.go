package usecases

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type ValidateOrderToDoneUseCase struct {
	repository ports.OrderRepository
}

func NewValidateOrderToDoneUseCase(repository ports.OrderRepository) *ValidateOrderToDoneUseCase {
	return &ValidateOrderToDoneUseCase{
		repository: repository,
	}
}

func (usecase *ValidateOrderToDoneUseCase) Execute(ctx context.Context, orderId uint) error {
	response, err := usecase.repository.GetOrderById(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "ValidateOrderToDoneUseCase -> GetOrderById")
	}

	if response.OrderStatus != "Preparando" {
		return &responses.BusinessResponse{
			StatusCode: 428,
			Message:    "The order must be in Preparando status",
		}
	}

	return nil
}

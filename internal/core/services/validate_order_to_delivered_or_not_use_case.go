package services

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type ValidateOrderToDeliveredOrNotUseCase struct {
	repository ports.OrderRepository
}

func NewValidateOrderToDeliveredOrNotUseCase(repository ports.OrderRepository) *ValidateOrderToDeliveredOrNotUseCase {
	return &ValidateOrderToDeliveredOrNotUseCase{
		repository: repository,
	}
}

func (usecase *ValidateOrderToDeliveredOrNotUseCase) Execute(ctx context.Context, orderId uint) error {
	response, err := usecase.repository.GetOrderById(ctx, orderId)

	if err != nil {
		return responses.GetResponseError(err, "ValidateOrderToDoneUseCase -> GetOrderById")
	}

	if response.OrderStatus != "Finalizado" {
		return &responses.BusinessResponse{
			StatusCode: 428,
			Message:    "The order must be in Finalizado status",
		}
	}

	return nil
}

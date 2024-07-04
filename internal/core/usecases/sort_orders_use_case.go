package usecases

import (
	"slices"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

type SortOrdersUseCase struct{}

func NewSortOrdersUseCase() *SortOrdersUseCase {
	return &SortOrdersUseCase{}
}

func (usecase *SortOrdersUseCase) Execute(orders []domain.OrderResponse) {
	slices.SortFunc(orders, func(previous, next domain.OrderResponse) int {
		if next.OrderStatus == "Finalizado" && (previous.OrderStatus == "Preparando" || previous.OrderStatus == "Criado") {
			return 1
		}

		if next.OrderStatus == "Preparando" && previous.OrderStatus == "Criado" {
			return 0
		}

		return -1
	})
}

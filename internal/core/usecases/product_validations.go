package usecases

import "github.com/thiagoluis88git/tech1/internal/core/domain"

type ValidateProductCategoryUseCase struct{}

func NewValidateProductCategoryUseCase() *ValidateProductCategoryUseCase {
	return &ValidateProductCategoryUseCase{}
}

func (usecase *ValidateProductCategoryUseCase) Execute(product domain.ProductForm) bool {
	if product.Category == "Combo" {
		return product.ComboProductsIds != nil && len(*product.ComboProductsIds) > 0
	}

	return true
}

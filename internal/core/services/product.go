package services

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type ProductService struct {
	repository ports.ProductRepository
}

func NewProductService(repository ports.ProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (service *ProductService) CreateProduct(ctx context.Context, product domain.ProductForm) (uint, error) {
	productId, err := service.repository.CreateProduct(ctx, product)

	if err != nil {
		return 0, responses.GetResponseError(err, "ProductService")
	}

	return productId, nil
}

func (service *ProductService) CreateCombo(ctx context.Context, combo domain.ComboForm) (uint, error) {
	productId, err := service.repository.CreateCombo(ctx, combo)

	if err != nil {
		return 0, responses.GetResponseError(err, "ProductService")
	}

	return productId, nil
}

func (service *ProductService) GetProductsByCategory(ctx context.Context, category string) ([]domain.ProductForm, error) {
	products, err := service.repository.GetProductsByCategory(ctx, category)

	if err != nil {
		return []domain.ProductForm{}, responses.GetResponseError(err, "ProductService")
	}

	return products, nil
}

func (service *ProductService) GetCombos(ctx context.Context) ([]domain.Combo, error) {
	combos, err := service.repository.GetCombos(ctx)

	if err != nil {
		return []domain.Combo{}, responses.GetResponseError(err, "ProductService")
	}

	return combos, nil
}

func (service *ProductService) DeleteProduct(ctx context.Context, productId uint) error {
	err := service.repository.DeleteProduct(ctx, productId)

	if err != nil {
		return responses.GetResponseError(err, "ProductService")
	}

	return nil
}

func (service *ProductService) UpdateProduct(ctx context.Context, product domain.ProductForm) error {
	err := service.repository.UpdateProduct(ctx, product)

	if err != nil {
		return responses.GetResponseError(err, "ProductService")
	}

	return nil
}

func (service *ProductService) DeleteCombo(ctx context.Context, comboId uint) error {
	err := service.repository.DeleteCombo(ctx, comboId)

	if err != nil {
		return responses.GetResponseError(err, "ProductService")
	}

	return nil
}

func (service *ProductService) UpdateCombo(ctx context.Context, combo domain.ComboForm) error {
	err := service.repository.UpdateCombo(ctx, combo)

	if err != nil {
		return responses.GetResponseError(err, "ProductService")
	}

	return nil
}

func (service *ProductService) GetCategories() []string {
	return service.repository.GetCategories()
}

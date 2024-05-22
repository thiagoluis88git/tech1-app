package services

import (
	"context"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type ProductService struct {
	repository      ports.ProductRepository
	validateUseCase *ValidateProductCategoryUseCase
}

func NewProductService(validateUseCase *ValidateProductCategoryUseCase, repository ports.ProductRepository) *ProductService {
	return &ProductService{
		repository:      repository,
		validateUseCase: validateUseCase,
	}
}

func (service *ProductService) CreateProduct(ctx context.Context, product domain.ProductForm) (uint, error) {
	if !service.validateUseCase.Execute(product) {
		return 0, &responses.BusinessResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Combo needs products",
		}
	}

	productId, err := service.repository.CreateProduct(ctx, product)

	if err != nil {
		return 0, responses.GetResponseError(err, "ProductService")
	}

	return productId, nil
}

func (service *ProductService) GetProductsByCategory(ctx context.Context, category string) ([]domain.ProductResponse, error) {
	products, err := service.repository.GetProductsByCategory(ctx, category)

	if err != nil {
		return []domain.ProductResponse{}, responses.GetResponseError(err, "ProductService")
	}

	return products, nil
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

func (service *ProductService) GetCategories() []string {
	return service.repository.GetCategories()
}

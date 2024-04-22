package services

import (
	"context"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/ports"
	"thiagoluis88git/tech1/pkg/responses"
)

type ProductService struct {
	repository ports.ProductRepository
}

func NewProductService(repository ports.ProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (service *ProductService) CreateProduct(ctx context.Context, product domain.Product) (uint, error) {
	productId, err := service.repository.CreateProduct(ctx, product)

	if err != nil {
		return 0, responses.GetResponseError(err, "ProductService")
	}

	return productId, nil
}

func (service *ProductService) GetProductsByCategory(ctx context.Context, category string) ([]domain.Product, error) {
	products, err := service.repository.GetProductsByCategory(ctx, category)

	if err != nil {
		return []domain.Product{}, responses.GetResponseError(err, "ProductService")
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

func (service *ProductService) UpdateProduct(ctx context.Context, product domain.Product) error {
	err := service.repository.UpdateProduct(ctx, product)

	if err != nil {
		return responses.GetResponseError(err, "ProductService")
	}

	return nil
}

func (service *ProductService) GetCategories() []string {
	return service.repository.GetCategories()
}

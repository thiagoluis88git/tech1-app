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

func (service *ProductService) GetCategories() []string {
	return service.repository.GetCategories()
}

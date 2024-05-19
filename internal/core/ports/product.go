package ports

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product domain.ProductForm) (uint, error)
	GetCategories() []string
	GetProductsByCategory(ctx context.Context, category string) ([]domain.ProductResponse, error)
	DeleteProduct(ctx context.Context, productId uint) error
	UpdateProduct(ctx context.Context, product domain.ProductForm) error
}

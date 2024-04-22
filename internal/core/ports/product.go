package ports

import (
	"context"
	"thiagoluis88git/tech1/internal/core/domain"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product domain.Product) (uint, error)
	GetCategories() []string
	GetProductsByCategory(ctx context.Context, category string) ([]domain.Product, error)
	DeleteProduct(ctx context.Context, productId uint) error
	UpdateProduct(ctx context.Context, product domain.Product) error
}

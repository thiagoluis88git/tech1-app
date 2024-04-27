package ports

import (
	"context"
	"thiagoluis88git/tech1/internal/core/domain"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product domain.Product) (uint, error)
	CreateCombo(ctx context.Context, combo domain.ComboForm) (uint, error)
	GetCategories() []string
	GetProductsByCategory(ctx context.Context, category string) ([]domain.Product, error)
	GetCombos(ctx context.Context) ([]domain.Combo, error)
	DeleteProduct(ctx context.Context, productId uint) error
	UpdateProduct(ctx context.Context, product domain.Product) error
}

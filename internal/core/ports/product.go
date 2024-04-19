package ports

import (
	"context"
	"thiagoluis88git/tech1/internal/core/domain"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product domain.Product) (uint, error)
	GetCategories() []string
}

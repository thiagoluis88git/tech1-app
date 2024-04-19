package ports

import (
	"context"
	"thiagoluis88git/tech1/internal/core/domain"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer domain.Customer) (uint, error)
}

package ports

import (
	"context"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer domain.Customer) (uint, error)
	UpdateCustomer(ctx context.Context, customer domain.Customer) error
	GetCustomerByCPF(ctx context.Context, cpf string) (domain.Customer, error)
	GetCustomerById(ctx context.Context, id uint) (domain.Customer, error)
}

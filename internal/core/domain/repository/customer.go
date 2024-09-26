package repository

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer dto.Customer) (uint, error)
	UpdateCustomer(ctx context.Context, customer dto.Customer) error
	GetCustomerById(ctx context.Context, id uint) (dto.Customer, error)
	GetCustomerByCPF(ctx context.Context, cpf string) (dto.Customer, error)
	Login(ctx context.Context, cpf string) (string, error)
	LoginUnknown() (string, error)
}

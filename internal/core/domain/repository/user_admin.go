package repository

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
)

type UserAdminRepository interface {
	CreateUser(ctx context.Context, customer dto.UserAdmin) (uint, error)
	UpdateUser(ctx context.Context, customer dto.UserAdmin) error
	GetUserById(ctx context.Context, id uint) (dto.UserAdmin, error)
	GetUserByCPF(ctx context.Context, cpf string) (dto.UserAdmin, error)
	Login(ctx context.Context, cpf string) (string, error)
}

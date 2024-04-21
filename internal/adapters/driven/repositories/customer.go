package repositories

import (
	"context"
	"thiagoluis88git/tech1/internal/adapters/driven/entities"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/ports"
	"thiagoluis88git/tech1/pkg/responses"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) ports.CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (repository *CustomerRepository) CreateCustomer(ctx context.Context, customer domain.Customer) (uint, error) {
	customerEntity := &entities.Customer{
		Name:  customer.Name,
		CPF:   customer.CPF,
		Email: customer.Email,
	}

	err := repository.db.Create(customerEntity).Error

	if err != nil {
		return 0, responses.GetDatabaseError(err)
	}

	return customerEntity.ID, nil
}

func (repository *CustomerRepository) GetCustomerByCPF(ctx context.Context, cpf string) (domain.Customer, error) {
	var customerEntity entities.Customer
	err := repository.db.WithContext(ctx).Where("cpf = ?", cpf).First(&customerEntity).Error

	if err != nil {
		return domain.Customer{}, responses.GetDatabaseError(err)
	}

	return domain.Customer{
		ID:    customerEntity.ID,
		Name:  customerEntity.Name,
		CPF:   customerEntity.CPF,
		Email: customerEntity.Email,
	}, nil
}

package repositories

import (
	"context"
	"github.com/thiagoluis88git/tech1/internal/adapters/driven/entities"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"

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

	err := repository.db.WithContext(ctx).Create(customerEntity).Error

	if err != nil {
		return 0, responses.GetDatabaseError(err)
	}

	return customerEntity.ID, nil
}

func (repository *CustomerRepository) UpdateCustomer(ctx context.Context, customer domain.Customer) error {
	customerEntity := &entities.Customer{
		Model: gorm.Model{ID: customer.ID},
		Name:  customer.Name,
		CPF:   customer.CPF,
		Email: customer.Email,
	}

	err := repository.db.WithContext(ctx).Save(&customerEntity).Error

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}

func (repository *CustomerRepository) GetCustomerByCPF(ctx context.Context, cpf string) (domain.Customer, error) {
	var customerEntity entities.Customer
	err := repository.db.WithContext(ctx).Where("cpf = ?", cpf).First(&customerEntity).Error

	if err != nil {
		return domain.Customer{}, responses.GetDatabaseError(err)
	}

	return repository.populateCustomer(customerEntity), nil
}

// GetCustomerById implements ports.CustomerRepository.
func (repository *CustomerRepository) GetCustomerById(ctx context.Context, id uint) (domain.Customer, error) {
	var customerEntity entities.Customer

	err := repository.
		db.WithContext(ctx).
		First(&customerEntity, id).
		Error

	if err != nil {
		return domain.Customer{}, responses.GetDatabaseError(err)
	}

	return repository.populateCustomer(customerEntity), nil
}

func (repository *CustomerRepository) populateCustomer(customerEntity entities.Customer) domain.Customer {
	return domain.Customer{
		ID:    customerEntity.ID,
		Name:  customerEntity.Name,
		CPF:   customerEntity.CPF,
		Email: customerEntity.Email,
	}
}

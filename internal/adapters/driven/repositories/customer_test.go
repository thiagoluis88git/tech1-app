package repositories

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/thiagoluis88git/tech1/internal/adapters/driven/model"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

func TestCustomerRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestCreateCustomerWithSuccess() {
	// ensure that the postgres database is empty
	var customers []model.Customer
	result := suite.db.Find(&customers)
	suite.NoError(result.Error)
	suite.Empty(customers)

	repo := NewCustomerRepository(suite.db)

	// Product 1
	newCustomer := domain.Customer{
		Name:  "Teste",
		CPF:   "12312312312",
		Email: "teste@teste.com",
	}

	newId, err := repo.CreateCustomer(suite.ctx, newCustomer)

	suite.NoError(err)
	suite.Equal(uint(1), newId)

	customer, err := repo.GetCustomerById(suite.ctx, uint(1))
	suite.NoError(err)
	suite.Equal(uint(1), customer.ID)
}

func (suite *RepositoryTestSuite) TestGetCustomerByCPFWithSuccess() {
	// ensure that the postgres database is empty
	var customers []model.Customer
	result := suite.db.Find(&customers)
	suite.NoError(result.Error)
	suite.Empty(customers)

	repo := NewCustomerRepository(suite.db)

	// Product 1
	newCustomer := domain.Customer{
		Name:  "Teste",
		CPF:   "12312312312",
		Email: "teste@teste.com",
	}

	newId, err := repo.CreateCustomer(suite.ctx, newCustomer)

	suite.NoError(err)
	suite.Equal(uint(1), newId)

	customer, err := repo.GetCustomerByCPF(suite.ctx, "12312312312")
	suite.NoError(err)
	suite.Equal(uint(1), customer.ID)
}

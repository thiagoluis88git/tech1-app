package repositories

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/thiagoluis88git/tech1/internal/core/data/model"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
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

	mockCognito := new(MockCognitoRemoteDataSource)
	repo := NewCustomerRepository(suite.db, mockCognito)

	newCustomer := dto.Customer{
		Name:  "Teste",
		CPF:   "12312312312",
		Email: "teste@teste.com",
	}

	newCustomerModel := &model.Customer{
		Name:  "Teste",
		CPF:   "12312312312",
		Email: "teste@teste.com",
	}

	mockCognito.On("SignUp", newCustomerModel).Return(nil)

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

	mockCognito := new(MockCognitoRemoteDataSource)
	repo := NewCustomerRepository(suite.db, mockCognito)

	// Product 1
	newCustomer := dto.Customer{
		Name:  "Teste",
		CPF:   "12312312312",
		Email: "teste@teste.com",
	}

	newCustomerModel := &model.Customer{
		Name:  "Teste",
		CPF:   "12312312312",
		Email: "teste@teste.com",
	}

	mockCognito.On("SignUp", newCustomerModel).Return(nil)

	newId, err := repo.CreateCustomer(suite.ctx, newCustomer)

	suite.NoError(err)
	suite.Equal(uint(1), newId)

	customer, err := repo.GetCustomerByCPF(suite.ctx, "12312312312")
	suite.NoError(err)
	suite.Equal(uint(1), customer.ID)
}

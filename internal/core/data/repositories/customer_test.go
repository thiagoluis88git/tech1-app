package repositories_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/thiagoluis88git/tech1/internal/core/data/model"
	"github.com/thiagoluis88git/tech1/internal/core/data/repositories"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func TestCustomerRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestCreateCustomerWithSuccess() {
	// ensure that the postgres database is empty
	var customers []model.Customer
	result := suite.db.Connection.Find(&customers)
	suite.NoError(result.Error)
	suite.Empty(customers)

	mockCognito := new(MockCognitoRemoteDataSource)
	repo := repositories.NewCustomerRepository(suite.db, mockCognito)

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

func (suite *RepositoryTestSuite) TestCreateCustomerWithConflictError() {
	// ensure that the postgres database is empty
	var customers []model.Customer
	result := suite.db.Connection.Find(&customers)
	suite.NoError(result.Error)
	suite.Empty(customers)

	mockCognito := new(MockCognitoRemoteDataSource)
	repo := repositories.NewCustomerRepository(suite.db, mockCognito)

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

	newIdError, err := repo.CreateCustomer(suite.ctx, newCustomer)

	suite.Error(err)
	suite.Equal(uint(0), newIdError)
}

func (suite *RepositoryTestSuite) TestCreateCustomerWithSignupError() {
	// ensure that the postgres database is empty
	var customers []model.Customer
	result := suite.db.Connection.Find(&customers)
	suite.NoError(result.Error)
	suite.Empty(customers)

	mockCognito := new(MockCognitoRemoteDataSource)
	repo := repositories.NewCustomerRepository(suite.db, mockCognito)

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

	mockCognito.On("SignUp", newCustomerModel).Return(&responses.NetworkError{
		Code: 419,
	})

	newId, err := repo.CreateCustomer(suite.ctx, newCustomer)

	suite.Error(err)
	suite.Equal(uint(0), newId)
}

func (suite *RepositoryTestSuite) TestUpdateCustomerWithSuccess() {
	// ensure that the postgres database is empty
	var customers []model.Customer
	result := suite.db.Connection.Find(&customers)
	suite.NoError(result.Error)
	suite.Empty(customers)

	mockCognito := new(MockCognitoRemoteDataSource)
	repo := repositories.NewCustomerRepository(suite.db, mockCognito)

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

	updateCustomerModel := dto.Customer{
		ID:    newId,
		Name:  "Teste 2",
		CPF:   "12312312312",
		Email: "teste@teste.com",
	}

	err = repo.UpdateCustomer(suite.ctx, updateCustomerModel)

	suite.NoError(err)

	customer, err := repo.GetCustomerById(suite.ctx, uint(1))
	suite.NoError(err)
	suite.Equal("Teste 2", customer.Name)
}

func (suite *RepositoryTestSuite) TestGetCustomerByCPFWithSuccess() {
	// ensure that the postgres database is empty
	var customers []model.Customer
	result := suite.db.Connection.Find(&customers)
	suite.NoError(result.Error)
	suite.Empty(customers)

	mockCognito := new(MockCognitoRemoteDataSource)
	repo := repositories.NewCustomerRepository(suite.db, mockCognito)

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

func (suite *RepositoryTestSuite) TestGetCustomerByCPFWithUnknownCustomerError() {
	// ensure that the postgres database is empty
	var customers []model.Customer
	result := suite.db.Connection.Find(&customers)
	suite.NoError(result.Error)
	suite.Empty(customers)

	mockCognito := new(MockCognitoRemoteDataSource)
	repo := repositories.NewCustomerRepository(suite.db, mockCognito)

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

	customer, err := repo.GetCustomerByCPF(suite.ctx, "12312")
	suite.Error(err)
	suite.Empty(customer)
}

func (suite *RepositoryTestSuite) TestLoginWithSuccess() {
	mockCognito := new(MockCognitoRemoteDataSource)
	repo := repositories.NewCustomerRepository(suite.db, mockCognito)

	mockCognito.On("Login", "123456").Return("TOKEN", nil)

	token, err := repo.Login(context.TODO(), "123456")

	suite.NoError(err)
	suite.Equal("TOKEN", token)
}

func (suite *RepositoryTestSuite) TestLoginWithCognitoError() {
	mockCognito := new(MockCognitoRemoteDataSource)
	repo := repositories.NewCustomerRepository(suite.db, mockCognito)

	mockCognito.On("Login", "123456").Return("", &responses.NetworkError{
		Code: 401,
	})

	token, err := repo.Login(context.TODO(), "123456")

	suite.Error(err)
	suite.Empty(token)
}

func (suite *RepositoryTestSuite) TestLoginUnknownWithSuccess() {
	mockCognito := new(MockCognitoRemoteDataSource)
	repo := repositories.NewCustomerRepository(suite.db, mockCognito)

	mockCognito.On("LoginUnknown").Return("TOKEN", nil)

	token, err := repo.LoginUnknown()

	suite.NoError(err)
	suite.Equal("TOKEN", token)
}

func (suite *RepositoryTestSuite) TestLoginUnknownWithCognitoError() {
	mockCognito := new(MockCognitoRemoteDataSource)
	repo := repositories.NewCustomerRepository(suite.db, mockCognito)

	mockCognito.On("LoginUnknown").Return("", &responses.NetworkError{
		Code: 401,
	})

	token, err := repo.LoginUnknown()

	suite.Error(err)
	suite.Empty(token)
}

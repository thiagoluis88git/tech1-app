package repositories

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/thiagoluis88git/tech1/internal/core/data/model"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
)

func TestPaymentRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestCreatePaymentOrderWithSuccess() {
	// ensure that the postgres database is empty
	var payments []model.Payment
	result := suite.db.Find(&payments)
	suite.NoError(result.Error)
	suite.Empty(payments)

	repo := NewPaymentRepository(suite.db)

	newPayment := dto.Payment{
		TotalPrice:  58.90,
		PaymentType: "Crédito",
	}

	payment, err := repo.CreatePaymentOrder(suite.ctx, newPayment)

	suite.NoError(err)
	suite.Equal(uint(1), payment.PaymentId)
}

func (suite *RepositoryTestSuite) TestCreatePaymentOrderWithUnknownUserError() {
	// ensure that the postgres database is empty
	var payments []model.Payment
	result := suite.db.Find(&payments)
	suite.NoError(result.Error)
	suite.Empty(payments)

	repo := NewPaymentRepository(suite.db)

	unknownUser := uint(2)

	newPayment := dto.Payment{
		CustomerID:  &unknownUser,
		TotalPrice:  58.90,
		PaymentType: "Crédito",
	}

	payment, err := repo.CreatePaymentOrder(suite.ctx, newPayment)

	suite.Error(err)
	suite.Empty(payment)
}

func (suite *RepositoryTestSuite) TestFinishPaymentOrderWithSuccess() {
	// ensure that the postgres database is empty
	var payments []model.Payment
	result := suite.db.Find(&payments)
	suite.NoError(result.Error)
	suite.Empty(payments)

	repo := NewPaymentRepository(suite.db)

	newPayment := dto.Payment{
		TotalPrice:  58.90,
		PaymentType: "Crédito",
	}

	payment, err := repo.CreatePaymentOrder(suite.ctx, newPayment)

	suite.NoError(err)
	suite.Equal(uint(1), payment.PaymentId)

	err = repo.FinishPaymentWithSuccess(suite.ctx, payment.PaymentId)

	suite.NoError(err)
}

func (suite *RepositoryTestSuite) TestFinishPaymentOrderWithError() {
	// ensure that the postgres database is empty
	var payments []model.Payment
	result := suite.db.Find(&payments)
	suite.NoError(result.Error)
	suite.Empty(payments)

	repo := NewPaymentRepository(suite.db)

	newPayment := dto.Payment{
		TotalPrice:  58.90,
		PaymentType: "Crédito",
	}

	payment, err := repo.CreatePaymentOrder(suite.ctx, newPayment)

	suite.NoError(err)
	suite.Equal(uint(1), payment.PaymentId)

	err = repo.FinishPaymentWithError(suite.ctx, payment.PaymentId)

	suite.NoError(err)
}

func (suite *RepositoryTestSuite) TestGetPaymentTypesWithSuccess() {
	repo := NewPaymentRepository(suite.db)

	paymentTypes := repo.GetPaymentTypes()

	suite.Equal("QR Code (Mercado Pago)", paymentTypes[0])
	suite.Equal("Crédito", paymentTypes[1])
}

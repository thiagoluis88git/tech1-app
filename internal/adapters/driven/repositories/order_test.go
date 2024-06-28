package repositories

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/thiagoluis88git/tech1/internal/adapters/driven/entities"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

func TestOrderRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestCreateOrderWithSuccess() {
	// ensure that the postgres database is empty
	var products []entities.Product
	result := suite.db.Find(&products)
	suite.NoError(result.Error)
	suite.Empty(products)

	repoProduct := NewProductRepository(suite.db)
	newProduct := domain.ProductForm{
		Name:        "New Product Created",
		Description: "New Description Product Created",
		Category:    "Category",
		Price:       2990,
		Images: []domain.ProducImage{
			{
				ImageUrl: "NewImageUrl",
			},
		},
	}

	newId, err := repoProduct.CreateProduct(suite.ctx, newProduct)
	suite.NoError(err)
	suite.Equal(uint(1), newId)

	repo := NewOrderRespository(suite.db)
	newOrder := domain.Order{
		TotalPrice:   5090,
		PaymentID:    uint(12),
		TicketNumber: 12,
		OrderProduct: []domain.OrderProduct{
			{
				ProductID: uint(1),
			},
		},
	}
	orderResponse, err := repo.CreateOrder(suite.ctx, newOrder)
	suite.NoError(err)
	suite.Equal(uint(1), orderResponse.OrderId)
}

package repositories

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/thiagoluis88git/tech1/internal/core/data/model"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
)

func TestOrderRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestCreateOrderWithSuccess() {
	// ensure that the postgres database is empty
	var products []model.Product
	result := suite.db.Find(&products)
	suite.NoError(result.Error)
	suite.Empty(products)

	repoProduct := NewProductRepository(suite.db)
	newProduct := dto.ProductForm{
		Name:        "New Product Created",
		Description: "New Description Product Created",
		Category:    "Category",
		Price:       2990,
		Images: []dto.ProducImage{
			{
				ImageUrl: "NewImageUrl",
			},
		},
	}

	newId, err := repoProduct.CreateProduct(suite.ctx, newProduct)
	suite.NoError(err)
	suite.Equal(uint(1), newId)

	repo := NewOrderRespository(suite.db)
	newOrder := dto.Order{
		TotalPrice:   5090,
		PaymentID:    uint(12),
		TicketNumber: 12,
		OrderProduct: []dto.OrderProduct{
			{
				ProductID: uint(1),
			},
		},
	}
	orderResponse, err := repo.CreateOrder(suite.ctx, newOrder)
	suite.NoError(err)
	suite.Equal(uint(1), orderResponse.OrderId)
}

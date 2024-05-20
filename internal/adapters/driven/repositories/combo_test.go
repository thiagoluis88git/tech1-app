package repositories

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/thiagoluis88git/tech1/internal/adapters/driven/entities"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
)

func TestComboRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestCreateComboWithSuccess() {
	// ensure that the postgres database is empty
	var products []entities.Product
	result := suite.db.Find(&products)
	suite.NoError(result.Error)
	suite.Empty(products)

	repo := NewProductRepository(suite.db)

	// Product 1
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
	newId, err := repo.CreateProduct(suite.ctx, newProduct)

	suite.NoError(err)
	suite.Equal(uint(1), newId)

	// Product 2
	newProduct2 := domain.ProductForm{
		Name:        "New Product Created 2",
		Description: "New Description Product Created 2",
		Category:    "Category",
		Price:       990,
		Images: []domain.ProducImage{
			{
				ImageUrl: "NewImageUrl 2",
			},
		},
	}
	newId2, err := repo.CreateProduct(suite.ctx, newProduct2)

	suite.NoError(err)
	suite.Equal(uint(2), newId2)

	// Product 3
	newProduct3 := domain.ProductForm{
		Name:        "New Product Created 3",
		Description: "New Description Product Created 3",
		Category:    "Category",
		Price:       1990,
		Images: []domain.ProducImage{
			{
				ImageUrl: "NewImageUrl",
			},
		},
	}
	newId3, err := repo.CreateProduct(suite.ctx, newProduct3)

	suite.NoError(err)
	suite.Equal(uint(3), newId3)

	// ensure that we have a new product in the database
	result = suite.db.Find(&products)
	suite.NoError(result.Error)
	suite.Equal(3, len(products))

	// Combo
	newCombo := domain.ProductForm{
		Name:        "New Combo",
		Description: "New Description Combo",
		Category:    "Combo",
		Price:       1990,
		Images: []domain.ProducImage{
			{
				ImageUrl: "NewImageUrl",
			},
		},
		ComboProductsIds: &[]uint{1, 2, 3},
	}

	comboId, err := repo.CreateProduct(suite.ctx, newCombo)
	suite.NoError(err)
	suite.Equal(uint(4), comboId)
}

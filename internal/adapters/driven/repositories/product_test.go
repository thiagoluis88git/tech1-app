package repositories

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/thiagoluis88git/tech1/internal/adapters/driven/entities"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func TestProductRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestGetProductsWithSuccess() {
	// ensure that the postgres database is empty
	var products []entities.Product
	result := suite.db.Find(&products)
	suite.NoError(result.Error)
	suite.Empty(products)

	repo := NewProductRepository(suite.db)

	newProduct := domain.Product{
		Name:        "New Product Created",
		Description: "New Description Product Created",
		Category:    "Lanches",
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

	createdProducts, err := repo.GetProductsByCategory(suite.ctx, "Lanches")

	suite.NoError(err)
	suite.Equal(1, len(createdProducts))
	suite.Equal("New Product Created", createdProducts[0].Name)
}

func (suite *RepositoryTestSuite) TestCreateProductWithSuccess() {
	// ensure that the postgres database is empty
	var products []entities.Product
	result := suite.db.Find(&products)
	suite.NoError(result.Error)
	suite.Empty(products)

	repo := NewProductRepository(suite.db)
	newProduct := domain.Product{
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

	// ensure that we have a new product in the database
	result = suite.db.Find(&products)
	suite.NoError(result.Error)
	suite.Equal(1, len(products))
	suite.Equal(uint(1), products[0].ID)
	suite.Equal("New Product Created", products[0].Name)
	suite.Equal("New Description Product Created", products[0].Description)
}

func (suite *RepositoryTestSuite) TestCreateComboWithSuccess() {
	// ensure that the postgres database is empty
	var combos []entities.Combo
	result := suite.db.Find(&combos)
	suite.NoError(result.Error)
	suite.Empty(combos)

	repo := NewProductRepository(suite.db)

	newProduct := domain.Product{
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

	newCombo := domain.ComboForm{
		Name:        "New Combo Created",
		Description: "New Description Combo Created",
		Price:       2990,
		Products:    []uint{1},
	}

	newId, err = repo.CreateCombo(suite.ctx, newCombo)
	suite.NoError(err)
	suite.Equal(uint(1), newId)

	newCombos, err := repo.GetCombos(suite.ctx)

	suite.NoError(err)
	suite.NoError(result.Error)
	suite.Equal(1, len(newCombos))
	suite.Equal(uint(1), newCombos[0].Id)
	suite.Equal("New Combo Created", newCombos[0].Name)
	suite.Equal("New Description Combo Created", newCombos[0].Description)
	suite.Equal(uint(1), newCombos[0].Products[0].Id)
}

func (suite *RepositoryTestSuite) TestUpdateComboWithSuccess() {
	// ensure that the postgres database is empty
	var combos []entities.Combo
	result := suite.db.Find(&combos)
	suite.NoError(result.Error)
	suite.Empty(combos)

	repo := NewProductRepository(suite.db)

	newProduct := domain.Product{
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

	newCombo := domain.ComboForm{
		Name:        "New Combo Created",
		Description: "New Description Combo Created",
		Price:       2990,
		Products:    []uint{1},
	}

	newId, err = repo.CreateCombo(suite.ctx, newCombo)
	suite.NoError(err)
	suite.Equal(uint(1), newId)

	updateCombo := domain.ComboForm{
		Id:          uint(1),
		Name:        "Updated Combo Created",
		Description: "Updated Description Combo Created",
		Price:       3990,
		Products:    []uint{1},
	}

	err = repo.UpdateCombo(suite.ctx, updateCombo)
	suite.NoError(err)

	newCombos, err := repo.GetCombos(suite.ctx)

	suite.NoError(err)
	suite.NoError(result.Error)
	suite.Equal(1, len(newCombos))
	suite.Equal(uint(1), newCombos[0].Id)
	suite.Equal("Updated Combo Created", newCombos[0].Name)
	suite.Equal("Updated Description Combo Created", newCombos[0].Description)
	suite.Equal(3990, newCombos[0].Price)
	suite.Equal(uint(1), newCombos[0].Products[0].Id)
}

func (suite *RepositoryTestSuite) TestUpdateProductWithSuccess() {
	// ensure that the postgres database is empty
	var products []entities.Product
	result := suite.db.Find(&products)
	suite.NoError(result.Error)
	suite.Empty(products)

	// create repository and save new note
	repo := NewProductRepository(suite.db)
	newProduct := domain.Product{
		Name:        "New Product",
		Description: "New Description Product",
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

	// ensure that we have a new product in the database
	result = suite.db.Find(&products)
	suite.NoError(result.Error)
	suite.Equal(1, len(products))
	suite.Equal(uint(1), products[0].ID)
	suite.Equal("New Product", products[0].Name)
	suite.Equal("New Description Product", products[0].Description)

	updateProduct := domain.Product{
		Id:          uint(1),
		Name:        "Updated Product",
		Description: "Updated Description Product",
		Category:    "Category",
		Price:       3990,
		Images: []domain.ProducImage{
			{
				ImageUrl: "NewImageUrl",
			},
		},
	}

	err = repo.UpdateProduct(suite.ctx, updateProduct)

	suite.NoError(err)

	// ensure that we have a new product in the database
	result = suite.db.Find(&products)
	suite.NoError(result.Error)
	suite.Equal(1, len(products))
	suite.Equal(uint(1), products[0].ID)
	suite.Equal("Updated Product", products[0].Name)
	suite.Equal("Updated Description Product", products[0].Description)
}

func (suite *RepositoryTestSuite) TestCreateProductWithConflictError() {
	// ensure that the postgres database is empty
	var products []entities.Product
	result := suite.db.Find(&products)
	suite.NoError(result.Error)
	suite.Empty(products)

	// create repository and save new note
	repo := NewProductRepository(suite.db)
	newProduct := domain.Product{
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

	newId, err = repo.CreateProduct(suite.ctx, newProduct)

	suite.Error(err)
	suite.Equal(uint(0), newId)

	var businessError *responses.LocalError
	suite.Equal(true, errors.As(err, &businessError))
	suite.Equal(responses.DATABASE_CONFLICT_ERROR, businessError.Code)
}

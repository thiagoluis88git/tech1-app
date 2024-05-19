package repositories

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/thiagoluis88git/tech1/internal/adapters/driven/entities"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/pkg/responses"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ProductRepositoryTestSuite struct {
	suite.Suite
	ctx                context.Context
	db                 *gorm.DB
	pgContainer        *postgres.PostgresContainer
	pgConnectionString string
}

func (suite *ProductRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := postgres.RunContainer(
		suite.ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithDatabase("notesdb"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	suite.NoError(err)

	connStr, err := pgContainer.ConnectionString(suite.ctx, "sslmode=disable")
	suite.NoError(err)

	db, err := gorm.Open(pg.Open(connStr), &gorm.Config{})
	suite.NoError(err)

	suite.pgContainer = pgContainer
	suite.pgConnectionString = connStr
	suite.db = db

	sqlDB, err := suite.db.DB()
	suite.NoError(err)

	err = sqlDB.Ping()
	suite.NoError(err)
}

func (suite *ProductRepositoryTestSuite) TearDownSuite() {
	err := suite.pgContainer.Terminate(suite.ctx)
	suite.NoError(err)
}

func (suite *ProductRepositoryTestSuite) SetupTest() {
	err := suite.db.AutoMigrate(
		&entities.Product{},
		&entities.ProductImage{},
		&entities.Combo{},
		&entities.ComboProduct{},
	)
	suite.NoError(err)
}

func (suite *ProductRepositoryTestSuite) TearDownTest() {
	suite.db.Exec("DROP TABLE IF EXISTS products CASCADE;")
	suite.db.Exec("DROP TABLE IF EXISTS product_images CASCADE;")
	suite.db.Exec("DROP TABLE IF EXISTS combos CASCADE;")
	suite.db.Exec("DROP TABLE IF EXISTS combo_products CASCADE;")
}

func (suite *ProductRepositoryTestSuite) TestCreateProductWithSuccess() {
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

	// ensure that we have a new product in the database
	result = suite.db.Find(&products)
	suite.NoError(result.Error)
	suite.Equal(1, len(products))
	suite.Equal(uint(1), products[0].ID)
	suite.Equal("New Product Created", products[0].Name)
	suite.Equal("New Description Product Created", products[0].Description)
}

func (suite *ProductRepositoryTestSuite) TestCreateProductWithConflictError() {
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

func TestProductServices(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}

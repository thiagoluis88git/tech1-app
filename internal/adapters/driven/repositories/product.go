package repositories

import (
	"context"
	"thiagoluis88git/tech1/internal/adapters/driven/entities"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/ports"
	"thiagoluis88git/tech1/pkg/responses"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ports.ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repository *ProductRepository) GetCategories() []string {
	return []string{
		entities.CategoryCombo,
		entities.CategorySnack,
		entities.CategoryBeverage,
		entities.CategoryToppings,
		entities.CategoryDesert,
	}
}

func (repository *ProductRepository) CreateProduct(ctx context.Context, product domain.Product) (uint, error) {
	tx := repository.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, responses.GetDatabaseError(err)
	}

	productEntity := &entities.Product{
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Price:       product.Price,
	}

	err := tx.Create(productEntity).Error

	if err != nil {
		tx.Rollback()
		return 0, responses.GetDatabaseError(err)
	}

	productImages := []*entities.ProductImage{}

	for _, value := range product.Images {
		productImages = append(productImages, &entities.ProductImage{
			ProductID: productEntity.ID,
			ImageUrl:  value.ImageUrl,
		})
	}

	err = tx.Create(productImages).Error

	if err != nil {
		tx.Rollback()
		return 0, responses.GetDatabaseError(err)
	}

	err = tx.Commit().Error

	if err != nil {
		tx.Rollback()
		return 0, responses.GetDatabaseError(err)
	}

	return productEntity.ID, nil
}

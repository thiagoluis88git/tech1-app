package repositories

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/adapters/driven/entities"
	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"

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
		domain.CategoryCombo,
		domain.CategorySnack,
		domain.CategoryBeverage,
		domain.CategoryToppings,
		domain.CategoryDesert,
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

func (repository *ProductRepository) CreateCombo(ctx context.Context, combo domain.ComboForm) (uint, error) {
	tx := repository.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, responses.GetDatabaseError(err)
	}

	comboEntity := &entities.Combo{
		Name:        combo.Name,
		Description: combo.Description,
		Price:       combo.Price,
	}

	err := tx.Create(comboEntity).Error

	if err != nil {
		tx.Rollback()
		return 0, responses.GetDatabaseError(err)
	}

	for _, value := range combo.Products {
		comboProductEntity := &entities.ComboProduct{
			ProductID: value,
			ComboID:   comboEntity.ID,
		}

		err := tx.Create(comboProductEntity).Error

		if err != nil {
			tx.Rollback()
			return 0, responses.GetDatabaseError(err)
		}
	}

	err = tx.Commit().Error

	if err != nil {
		tx.Rollback()
		return 0, responses.GetDatabaseError(err)
	}

	return comboEntity.ID, nil
}

func (repository *ProductRepository) GetProductsByCategory(ctx context.Context, category string) ([]domain.Product, error) {
	var productEntities []entities.Product
	err := repository.
		db.WithContext(ctx).
		Model(&entities.Product{}).
		Preload("ProductImage").
		Where("category = ?", category).
		Find(&productEntities).
		Error

	if err != nil {
		return []domain.Product{}, responses.GetDatabaseError(err)
	}

	return repository.buildProducts(productEntities), nil
}

func (repository *ProductRepository) GetCombos(ctx context.Context) ([]domain.Combo, error) {
	var comboEntities []entities.Combo
	err := repository.
		db.WithContext(ctx).
		Model(&entities.Combo{}).
		Preload("Products").
		Find(&comboEntities).
		Error

	if err != nil {
		return []domain.Combo{}, responses.GetDatabaseError(err)
	}

	combos := []domain.Combo{}

	for _, value := range comboEntities {
		combo := domain.Combo{
			Id:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			Price:       value.Price,
		}

		products := []domain.Product{}

		for _, value := range value.Products {
			var productEntities []entities.Product
			err := repository.
				db.WithContext(ctx).
				Model(&entities.Product{}).
				Preload("ProductImage").
				Where("id = ?", value.ProductID).
				Find(&productEntities).
				Limit(1).
				Error

			if err != nil {
				return []domain.Combo{}, responses.GetDatabaseError(err)
			}

			products = append(products, repository.buildProduct(productEntities[0]))
		}

		combo.Products = products
		combos = append(combos, combo)
	}

	return combos, nil
}

func (repository *ProductRepository) DeleteProduct(ctx context.Context, productId uint) error {
	tx := repository.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return responses.GetDatabaseError(err)
	}

	err := tx.Where("product_id = ?", productId).Delete(&entities.ProductImage{}).Error

	if err != nil {
		tx.Rollback()
		return responses.GetDatabaseError(err)
	}

	err = tx.Delete(&entities.Product{}, productId).Error

	if err != nil {
		tx.Rollback()
		return responses.GetDatabaseError(err)
	}

	err = tx.Commit().Error

	if err != nil {
		tx.Rollback()
		return responses.GetDatabaseError(err)
	}

	return nil
}

func (repository *ProductRepository) UpdateProduct(ctx context.Context, product domain.Product) error {
	productEntity := entities.Product{
		Model:       gorm.Model{ID: product.Id},
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Price:       product.Price,
	}

	err := repository.db.WithContext(ctx).Save(&productEntity).Error

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}

func (repository *ProductRepository) buildProducts(productEntities []entities.Product) []domain.Product {
	products := []domain.Product{}

	for _, value := range productEntities {
		products = append(products, repository.buildProduct(value))
	}

	return products
}

func (repository *ProductRepository) buildProduct(value entities.Product) domain.Product {
	images := []domain.ProducImage{}

	for _, valueImage := range value.ProductImage {
		images = append(images, domain.ProducImage{
			ImageUrl: valueImage.ImageUrl,
		})
	}

	return domain.Product{
		Id:          value.ID,
		Name:        value.Name,
		Description: value.Description,
		Category:    value.Category,
		Price:       value.Price,
		Images:      images,
	}
}

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
		entities.CategoryCombo,
		entities.CategorySnack,
		entities.CategoryBeverage,
		entities.CategoryToppings,
		entities.CategoryDesert,
	}
}

func (repository *ProductRepository) CreateProduct(ctx context.Context, product domain.ProductForm) (uint, error) {
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

	err = repository.createComboIfProductsNedded(tx, product, productEntity.ID)

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

func (repository *ProductRepository) createComboIfProductsNedded(
	tx *gorm.DB,
	productWithCombo domain.ProductForm,
	comboId uint,
) error {
	if productWithCombo.ComboProductsIds != nil {
		for _, value := range *productWithCombo.ComboProductsIds {
			comboProductEntity := &entities.ComboProduct{
				ProductID:      comboId,
				ComboProductID: value,
			}

			err := tx.Create(comboProductEntity).Error

			if err != nil {
				tx.Rollback()
				return responses.GetDatabaseError(err)
			}
		}
	}

	return nil
}

func (repository *ProductRepository) GetProductsByCategory(ctx context.Context, category string) ([]domain.ProductResponse, error) {
	var productEntities []entities.Product
	err := repository.
		db.WithContext(ctx).
		Model(&entities.Product{}).
		Preload("ProductImage").
		Preload("ComboProduct").
		Where("category = ?", category).
		Find(&productEntities).
		Error

	if err != nil {
		return []domain.ProductResponse{}, responses.GetDatabaseError(err)
	}

	return repository.buildProducts(ctx, productEntities), nil
}

func (repository *ProductRepository) GetProductById(ctx context.Context, id uint) (domain.ProductResponse, error) {
	var productEntity entities.Product
	err := repository.
		db.WithContext(ctx).
		Model(&entities.Product{}).
		Preload("ProductImage").
		Preload("ComboProduct").
		First(&productEntity, id).
		Error

	if err != nil {
		return domain.ProductResponse{}, responses.GetDatabaseError(err)
	}

	return repository.buildProduct(ctx, productEntity), nil
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

	err := tx.Where("product_id = ?", productId).Unscoped().Delete(&entities.ProductImage{}).Error

	if err != nil {
		tx.Rollback()
		return responses.GetDatabaseError(err)
	}

	err = tx.Where("product_id = ?", productId).Unscoped().Delete(&entities.ComboProduct{}).Error

	if err != nil {
		tx.Rollback()
		return responses.GetDatabaseError(err)
	}

	err = tx.Unscoped().Delete(&entities.Product{}, productId).Error

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

func (repository *ProductRepository) UpdateProduct(ctx context.Context, product domain.ProductForm) error {
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

func (repository *ProductRepository) buildProducts(ctx context.Context, productEntities []entities.Product) []domain.ProductResponse {
	products := []domain.ProductResponse{}

	for _, value := range productEntities {
		products = append(products, repository.buildProduct(ctx, value))
	}

	return products
}

func (repository *ProductRepository) buildProduct(ctx context.Context, value entities.Product) domain.ProductResponse {
	images := []domain.ProducImage{}

	for _, valueImage := range value.ProductImage {
		images = append(images, domain.ProducImage{
			ImageUrl: valueImage.ImageUrl,
		})
	}

	comboProducts := repository.getComboProductsIfNedded(ctx, value)

	return domain.ProductResponse{
		Id:            value.ID,
		Name:          value.Name,
		Description:   value.Description,
		Category:      value.Category,
		Price:         value.Price,
		Images:        images,
		ComboProducts: comboProducts,
	}
}

func (repository *ProductRepository) getComboProductsIfNedded(ctx context.Context, value entities.Product) *[]domain.ProductResponse {
	var comboProducts []domain.ProductResponse

	if value.ComboProduct != nil {
		comboProducts = make([]domain.ProductResponse, 0)

		for _, comboProduct := range value.ComboProduct {
			var product entities.Product

			err := repository.db.
				WithContext(ctx).
				Preload("ProductImage").
				First(&product, comboProduct.ComboProductID).
				Error

			if err == nil {
				comboProducts = append(comboProducts, repository.buildProduct(ctx, product))
			}
		}
	}

	return &comboProducts
}

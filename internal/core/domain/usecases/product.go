package usecases

import (
	"context"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/domain/repository"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type CreateProductUseCase interface {
	Execute(ctx context.Context, product dto.ProductForm) (uint, error)
}

type CreateProductUseCaseImpl struct {
	repository      repository.ProductRepository
	validateUseCase *ValidateProductCategoryUseCase
}

type GetProductsByCategoryUseCase interface {
	Execute(ctx context.Context, category string) ([]dto.ProductResponse, error)
}

type GetProductsByCategoryUseCaseImpl struct {
	repository repository.ProductRepository
}

type GetProductByIdUseCase interface {
	Execute(ctx context.Context, id uint) (dto.ProductResponse, error)
}

type GetProductByIdUseCaseImpl struct {
	repository repository.ProductRepository
}

type DeleteProductUseCase interface {
	Execute(ctx context.Context, productId uint) error
}

type DeleteProductUseCaseImpl struct {
	repository repository.ProductRepository
}

type UpdateProductUseCase interface {
	Execute(ctx context.Context, product dto.ProductForm) error
}

type UpdateProductUseCaseImpl struct {
	repository repository.ProductRepository
}

type GetCategoriesUseCase struct {
	repository repository.ProductRepository
}

func NewCreateProductUseCase(validateUseCase *ValidateProductCategoryUseCase, repository repository.ProductRepository) CreateProductUseCase {
	return &CreateProductUseCaseImpl{
		repository:      repository,
		validateUseCase: validateUseCase,
	}
}

func NewGetProductsByCategoryUseCase(repository repository.ProductRepository) GetProductsByCategoryUseCase {
	return &GetProductsByCategoryUseCaseImpl{
		repository: repository,
	}
}

func NewGetProductByIdUseCase(repository repository.ProductRepository) GetProductByIdUseCase {
	return &GetProductByIdUseCaseImpl{
		repository: repository,
	}
}

func NewDeleteProductUseCase(repository repository.ProductRepository) DeleteProductUseCase {
	return &DeleteProductUseCaseImpl{
		repository: repository,
	}
}

func NewUpdateProductUseCase(repository repository.ProductRepository) UpdateProductUseCase {
	return &UpdateProductUseCaseImpl{
		repository: repository,
	}
}

func NewGetCategoriesUseCase(repository repository.ProductRepository) *GetCategoriesUseCase {
	return &GetCategoriesUseCase{
		repository: repository,
	}
}

func (service *CreateProductUseCaseImpl) Execute(ctx context.Context, product dto.ProductForm) (uint, error) {
	if !service.validateUseCase.Execute(product) {
		return 0, &responses.BusinessResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Combo needs products",
		}
	}

	productId, err := service.repository.CreateProduct(ctx, product)

	if err != nil {
		return 0, responses.GetResponseError(err, "ProductService")
	}

	return productId, nil
}

func (service *GetProductsByCategoryUseCaseImpl) Execute(ctx context.Context, category string) ([]dto.ProductResponse, error) {
	products, err := service.repository.GetProductsByCategory(ctx, category)

	if err != nil {
		return []dto.ProductResponse{}, responses.GetResponseError(err, "ProductService")
	}

	return products, nil
}

func (service *GetProductByIdUseCaseImpl) Execute(ctx context.Context, id uint) (dto.ProductResponse, error) {
	products, err := service.repository.GetProductById(ctx, id)

	if err != nil {
		return dto.ProductResponse{}, responses.GetResponseError(err, "ProductService")
	}

	return products, nil
}

func (service *DeleteProductUseCaseImpl) Execute(ctx context.Context, productId uint) error {
	err := service.repository.DeleteProduct(ctx, productId)

	if err != nil {
		return responses.GetResponseError(err, "ProductService")
	}

	return nil
}

func (service *UpdateProductUseCaseImpl) Execute(ctx context.Context, product dto.ProductForm) error {
	err := service.repository.UpdateProduct(ctx, product)

	if err != nil {
		return responses.GetResponseError(err, "ProductService")
	}

	return nil
}

func (service *GetCategoriesUseCase) Execute() []string {
	return service.repository.GetCategories()
}

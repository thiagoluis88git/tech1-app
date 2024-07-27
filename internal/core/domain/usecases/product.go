package usecases

import (
	"context"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/domain/repository"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type CreateProductUseCase struct {
	repository      repository.ProductRepository
	validateUseCase *ValidateProductCategoryUseCase
}

type GetProductsByCategoryUseCase struct {
	repository repository.ProductRepository
}

type GetProductByIdUseCase struct {
	repository repository.ProductRepository
}

type DeleteProductUseCase struct {
	repository repository.ProductRepository
}

type UpdateProductUseCase struct {
	repository repository.ProductRepository
}

type GetCategoriesUseCase struct {
	repository repository.ProductRepository
}

func NewCreateProductUseCase(validateUseCase *ValidateProductCategoryUseCase, repository repository.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		repository:      repository,
		validateUseCase: validateUseCase,
	}
}

func NewGetProductsByCategoryUseCase(repository repository.ProductRepository) *GetProductsByCategoryUseCase {
	return &GetProductsByCategoryUseCase{
		repository: repository,
	}
}

func NewGetProductByIdUseCase(repository repository.ProductRepository) *GetProductByIdUseCase {
	return &GetProductByIdUseCase{
		repository: repository,
	}
}

func NewDeleteProductUseCase(repository repository.ProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		repository: repository,
	}
}

func NewUpdateProductUseCase(repository repository.ProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		repository: repository,
	}
}

func NewGetCategoriesUseCase(repository repository.ProductRepository) *GetCategoriesUseCase {
	return &GetCategoriesUseCase{
		repository: repository,
	}
}

func (service *CreateProductUseCase) Execute(ctx context.Context, product dto.ProductForm) (uint, error) {
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

func (service *GetProductsByCategoryUseCase) Execute(ctx context.Context, category string) ([]dto.ProductResponse, error) {
	products, err := service.repository.GetProductsByCategory(ctx, category)

	if err != nil {
		return []dto.ProductResponse{}, responses.GetResponseError(err, "ProductService")
	}

	return products, nil
}

func (service *GetProductByIdUseCase) Execute(ctx context.Context, id uint) (dto.ProductResponse, error) {
	products, err := service.repository.GetProductById(ctx, id)

	if err != nil {
		return dto.ProductResponse{}, responses.GetResponseError(err, "ProductService")
	}

	return products, nil
}

func (service *DeleteProductUseCase) Execute(ctx context.Context, productId uint) error {
	err := service.repository.DeleteProduct(ctx, productId)

	if err != nil {
		return responses.GetResponseError(err, "ProductService")
	}

	return nil
}

func (service *UpdateProductUseCase) Execute(ctx context.Context, product dto.ProductForm) error {
	err := service.repository.UpdateProduct(ctx, product)

	if err != nil {
		return responses.GetResponseError(err, "ProductService")
	}

	return nil
}

func (service *GetCategoriesUseCase) Execute() []string {
	return service.repository.GetCategories()
}

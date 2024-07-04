package usecases

import (
	"context"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/ports"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type CreateProductUseCase struct {
	repository      ports.ProductRepository
	validateUseCase *ValidateProductCategoryUseCase
}

type GetProductsByCategoryUseCase struct {
	repository ports.ProductRepository
}

type GetProductByIdUseCase struct {
	repository ports.ProductRepository
}

type DeleteProductUseCase struct {
	repository ports.ProductRepository
}

type UpdateProductUseCase struct {
	repository ports.ProductRepository
}

type GetCategoriesUseCase struct {
	repository ports.ProductRepository
}

func NewCreateProductUseCase(validateUseCase *ValidateProductCategoryUseCase, repository ports.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		repository:      repository,
		validateUseCase: validateUseCase,
	}
}

func NewGetProductsByCategoryUseCase(repository ports.ProductRepository) *GetProductsByCategoryUseCase {
	return &GetProductsByCategoryUseCase{
		repository: repository,
	}
}

func NewGetProductByIdUseCase(repository ports.ProductRepository) *GetProductByIdUseCase {
	return &GetProductByIdUseCase{
		repository: repository,
	}
}

func NewDeleteProductUseCase(repository ports.ProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		repository: repository,
	}
}

func NewUpdateProductUseCase(repository ports.ProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		repository: repository,
	}
}

func NewGetCategoriesUseCase(repository ports.ProductRepository) *GetCategoriesUseCase {
	return &GetCategoriesUseCase{
		repository: repository,
	}
}

func (service *CreateProductUseCase) Execute(ctx context.Context, product domain.ProductForm) (uint, error) {
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

func (service *GetProductsByCategoryUseCase) Execute(ctx context.Context, category string) ([]domain.ProductResponse, error) {
	products, err := service.repository.GetProductsByCategory(ctx, category)

	if err != nil {
		return []domain.ProductResponse{}, responses.GetResponseError(err, "ProductService")
	}

	return products, nil
}

func (service *GetProductByIdUseCase) Execute(ctx context.Context, id uint) (domain.ProductResponse, error) {
	products, err := service.repository.GetProductById(ctx, id)

	if err != nil {
		return domain.ProductResponse{}, responses.GetResponseError(err, "ProductService")
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

func (service *UpdateProductUseCase) Execute(ctx context.Context, product domain.ProductForm) error {
	err := service.repository.UpdateProduct(ctx, product)

	if err != nil {
		return responses.GetResponseError(err, "ProductService")
	}

	return nil
}

func (service *GetCategoriesUseCase) Execute() []string {
	return service.repository.GetCategories()
}

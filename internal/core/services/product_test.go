package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func TestProductServices(t *testing.T) {
	t.Run("got success when getting product categories in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockProductRepository)
		sut := NewProductService(mockRepo)

		mockRepo.On("GetCategories").Return([]string{"Combo", "Bebidas", "Lanches"})

		response := sut.GetCategories()

		assert.NotEmpty(t, response)

		assert.Equal(t, 3, len(response))
	})

	t.Run("got success when getting products by category in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockProductRepository)
		sut := NewProductService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("GetProductsByCategory", ctx, "category").Return(productsByCategory, nil)

		response, err := sut.GetProductsByCategory(ctx, "category")

		mockRepo.AssertExpectations(t)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		assert.Equal(t, 3, len(response))
	})

	t.Run("got error when getting products by category in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockProductRepository)
		sut := NewProductService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("GetProductsByCategory", ctx, "category").Return(uint(0), &responses.LocalError{
			Code:    3,
			Message: "DATABASE_CONFLICT_ERROR",
		})

		response, err := sut.GetProductsByCategory(ctx, "category")

		mockRepo.AssertExpectations(t)

		assert.Error(t, err)
		assert.Empty(t, response)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusConflict, businessError.StatusCode)
	})

	t.Run("got success when creating product in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockProductRepository)
		sut := NewProductService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("CreateProduct", ctx, productCreation).Return(uint(1), nil)

		response, err := sut.CreateProduct(ctx, productCreation)

		mockRepo.AssertExpectations(t)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		assert.Equal(t, uint(1), response)
	})

	t.Run("got error when creating product in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockProductRepository)
		sut := NewProductService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("CreateProduct", ctx, productCreation).Return(uint(0), &responses.LocalError{
			Code:    3,
			Message: "DATABASE_CONFLICT_ERROR",
		})

		response, err := sut.CreateProduct(ctx, productCreation)

		mockRepo.AssertExpectations(t)

		assert.Error(t, err)
		assert.Empty(t, response)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusConflict, businessError.StatusCode)
	})

	t.Run("got success when deleting product in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockProductRepository)
		sut := NewProductService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("DeleteProduct", ctx, uint(12)).Return(nil)

		err := sut.DeleteProduct(ctx, uint(12))

		mockRepo.AssertExpectations(t)

		assert.NoError(t, err)
	})

	t.Run("got error when deleting product in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockProductRepository)
		sut := NewProductService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("DeleteProduct", ctx, uint(12)).Return(&responses.LocalError{
			Code:    3,
			Message: "DATABASE_CONFLICT_ERROR",
		})

		err := sut.DeleteProduct(ctx, uint(12))

		mockRepo.AssertExpectations(t)

		assert.Error(t, err)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusConflict, businessError.StatusCode)
	})

	t.Run("got success when updating product in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockProductRepository)
		sut := NewProductService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("UpdateProduct", ctx, productUpdate).Return(nil)

		err := sut.UpdateProduct(ctx, productUpdate)

		mockRepo.AssertExpectations(t)

		assert.NoError(t, err)
	})

	t.Run("got error when updating product in services", func(t *testing.T) {
		t.Parallel()

		mockRepo := new(MockProductRepository)
		sut := NewProductService(mockRepo)

		ctx := context.TODO()

		mockRepo.On("UpdateProduct", ctx, productUpdate).Return(&responses.LocalError{
			Code:    3,
			Message: "DATABASE_CONFLICT_ERROR",
		})

		err := sut.UpdateProduct(ctx, productUpdate)

		mockRepo.AssertExpectations(t)

		assert.Error(t, err)

		var businessError *responses.BusinessResponse
		assert.Equal(t, true, errors.As(err, &businessError))
		assert.Equal(t, http.StatusConflict, businessError.StatusCode)
	})
}

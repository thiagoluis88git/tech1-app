package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/services"
	"github.com/thiagoluis88git/tech1/pkg/httpserver"
)

// @Summary Create new product
// @Description Create new product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body domain.ProductForm true "product"
// @Success 200 {object} domain.ProductResponse
// @Failure 400 "Product has required fields"
// @Failure 409 "This Product is already added"
// @Router /api/products [post]
func CreateProductHandler(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product domain.ProductForm

		err := httpserver.DecodeJSONBody(w, r, &product)

		if err != nil {
			log.Print("decoding product body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		productId, err := productService.CreateProduct(context.Background(), product)

		if err != nil {
			log.Print("create product", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, domain.ProductCreationResponse{
			Id: productId,
		})
	}
}

// @Summary List all products by a category
// @Description List all products by a category
// @Tags Product
// @Param category path string true "Lanches"
// @Accept json
// @Produce json
// @Success 200 {object} []domain.ProductResponse
// @Router /api/products/categories/{category} [get]
func GetProductsByCategoryHandler(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category, err := httpserver.GetPathParamFromRequest(r, "category")

		if err != nil {
			log.Print("get products by category", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		products, err := productService.GetProductsByCategory(context.Background(), category)

		if err != nil {
			log.Print("get products by category", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, products)
	}
}

// @Summary Get product by ID
// @Description Get product by ID
// @Tags Product
// @Param id path int true "12"
// @Accept json
// @Produce json
// @Success 200 {object} domain.ProductResponse
// @Router /api/products/{id} [get]
func GetProductsByIdHandler(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productIdStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("get product by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		productId, err := strconv.Atoi(productIdStr)

		if err != nil {
			log.Print("get product by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		product, err := productService.GetProductById(context.Background(), uint(productId))

		if err != nil {
			log.Print("get product by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, product)
	}
}

// @Summary Delete a product
// @Description Delete a product by ID
// @Tags Product
// @Param id path int true "12"
// @Accept json
// @Produce json
// @Success 204
// @Router /api/products/{id} [delete]
func DeleteProductHandler(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productIdStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("delete product", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		productId, err := strconv.Atoi(productIdStr)

		if err != nil {
			log.Print("delete product", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		err = productService.DeleteProduct(context.Background(), uint(productId))

		if err != nil {
			log.Print("delete product", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseNoContentSuccess(w)
	}
}

// @Summary Update a product
// @Description Update a product by ID
// @Tags Product
// @Param id path int true "12"
// @Accept json
// @Produce json
// @Success 204
// @Router /api/products/{id} [put]
func UpdateProductHandler(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productIdStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("update product", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		productId, err := strconv.Atoi(productIdStr)

		if err != nil {
			log.Print("update product", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		var product domain.ProductForm

		err = httpserver.DecodeJSONBody(w, r, &product)

		if err != nil {
			log.Print("decoding product body for update product", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		product.Id = uint(productId)
		err = productService.UpdateProduct(context.Background(), product)

		if err != nil {
			log.Print("update product", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseNoContentSuccess(w)
	}
}

// @Summary Get all categories
// @Description Get all categories to filter in products by category
// @Tags Product
// @Param id path int true "12"
// @Accept json
// @Produce json
// @Success 200 {object} []string
// @Router /api/products/categories [get]
func GetCategoriesHandler(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, productService.GetCategories())
	}
}

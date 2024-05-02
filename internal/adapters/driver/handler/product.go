package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/services"
	"thiagoluis88git/tech1/pkg/httpserver"
)

// @Summary Create new product
// @Description Create new product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body domain.Product true "product"
// @Success 200 {object} domain.ProductResponse
// @Failure 400 "Product has required fields"
// @Failure 409 "This Product is already added"
// @Router /api/products [post]
func CreateProductHandler(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product domain.Product

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

		httpserver.SendResponseSuccess(w, domain.ProductResponse{
			Id: productId,
		})
	}
}

// @Summary Create new combo
// @Description Create new combo of products
// @Tags Product
// @Accept json
// @Produce json
// @Param product body domain.ComboForm true "combo"
// @Success 200 {object} domain.ProductResponse
// @Failure 400 "ComboForm has required fields"
// @Failure 409 "This Combo is already added"
// @Router /api/products/combo [post]
func CreateComboHandler(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var combo domain.ComboForm

		err := httpserver.DecodeJSONBody(w, r, &combo)

		if err != nil {
			log.Print("decoding combo body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		productId, err := productService.CreateCombo(context.Background(), combo)

		if err != nil {
			log.Print("create combo", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, domain.ProductResponse{
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
// @Success 200 {object} []domain.Product
// @Router /api/products/{category} [get]
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

// @Summary List all combos
// @Description List all combos with their products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Combo
// @Router /api/products/combo [get]
func GetCombosHandler(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		combos, err := productService.GetCombos(context.Background())

		if err != nil {
			log.Print("get combos", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, combos)
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

		var product domain.Product

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

func GetCategoryHandler(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, productService.GetCategories())
	}
}

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

		httpserver.SendResponseSuccess(w, productId)
	}
}

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

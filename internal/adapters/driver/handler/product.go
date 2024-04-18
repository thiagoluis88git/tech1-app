package handler

import (
	"context"
	"log"
	"net/http"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/services"
	"thiagoluis88git/tech1/pkg/httpserver"
)

func CreateProductHandler(productService *services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product domain.Product

		err := httpserver.DecodeJSONBody(w, r, &product)

		if err != nil {
			log.Print("decoding invest fund body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		productId, err := productService.CreateProduct(context.Background(), product)

		if err != nil {
			log.Print("listing funds", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, productId)
	}
}

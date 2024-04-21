package handler

import (
	"context"
	"log"
	"net/http"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/services"
	"thiagoluis88git/tech1/pkg/httpserver"
)

func CreateOrderHandler(orderService *services.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order domain.Order

		err := httpserver.DecodeJSONBody(w, r, &order)

		if err != nil {
			log.Print("decoding order body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		response, err := orderService.CreateOrder(context.Background(), order)

		if err != nil {
			log.Print("create order", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, response)
	}
}

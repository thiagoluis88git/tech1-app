package handler

import (
	"context"
	"log"
	"net/http"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/services"
	"thiagoluis88git/tech1/pkg/httpserver"
)

func CreateCustomerHandler(customerService *services.CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer domain.Customer

		err := httpserver.DecodeJSONBody(w, r, &customer)

		if err != nil {
			log.Print("decoding customer body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		customerId, err := customerService.CreateCustomer(context.Background(), customer)

		if err != nil {
			log.Print("create customer", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, customerId)
	}
}

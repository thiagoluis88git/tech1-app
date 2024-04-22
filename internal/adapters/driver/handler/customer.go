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

func UpdateCustomerHandler(customerService *services.CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerIdStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("update customer", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		customerId, err := strconv.Atoi(customerIdStr)

		if err != nil {
			log.Print("update customer", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		var customer domain.Customer
		err = httpserver.DecodeJSONBody(w, r, &customer)

		if err != nil {
			log.Print("decoding customer body for update", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		customer.ID = uint(customerId)
		err = customerService.UpdateCustomer(context.Background(), customer)

		if err != nil {
			log.Print("update customer", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, customerId)
	}
}

func GetCustomerByCPFHandler(customerService *services.CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cpf, err := httpserver.GetPathParamFromRequest(r, "cpf")

		if err != nil {
			log.Print("get customer by cpf", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		customer, err := customerService.GetCustomerByCPF(context.Background(), cpf)

		if err != nil {
			log.Print("get customer by cpf", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, customer)
	}
}

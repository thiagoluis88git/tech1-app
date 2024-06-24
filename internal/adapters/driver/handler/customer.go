package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/services"
	"github.com/thiagoluis88git/tech1/pkg/httpserver"
)

// @Summary Create new customer
// @Description Create new customer. This process is not required to make an order
// @Tags Customer
// @Accept json
// @Produce json
// @Param product body domain.Customer true "customer"
// @Success 200 {object} domain.CustomerResponse
// @Failure 400 "Customer has required fields"
// @Failure 409 "This Customer is already added"
// @Router /api/customers [post]
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

		response, err := customerService.CreateCustomer(r.Context(), customer)

		if err != nil {
			log.Print("create customer", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, response)
	}
}

// @Summary Update customer
// @Description Update customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path int true "12"
// @Param product body domain.Customer true "customer"
// @Success 204
// @Failure 400 "Customer has required fields"
// @Failure 404 "Customer not found"
// @Router /api/customers/{id} [put]
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
		err = customerService.UpdateCustomer(r.Context(), customer)

		if err != nil {
			log.Print("update customer", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseNoContentSuccess(w)
	}
}

// @Summary Get customer by ID
// @Description Get customer by ID
// @Tags Customer
// @Accept json
// @Produce json
// @Param Id path string true "12"
// @Success 200 {object} domain.Customer
// @Failure 404 "Customer not found"
// @Router /api/customers/{id} [get]
func GetCustomerByIdHandler(customerService *services.CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerIdStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("get customer by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		customerId, err := strconv.Atoi(customerIdStr)

		if err != nil {
			log.Print("get customer by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		customer, err := customerService.GetCustomerById(r.Context(), uint(customerId))

		if err != nil {
			log.Print("get customer by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, customer)
	}
}

// @Summary Get customer by CPF
// @Description Get customer by CPF. This Endpoint can be used as a Login
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body domain.CustomerForm true "customerForm"
// @Success 200 {object} domain.Customer
// @Failure 404 "Customer not found"
// @Router /api/customers/login [post]
func GetCustomerByCPFHandler(customerService *services.CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customerForm domain.CustomerForm

		err := httpserver.DecodeJSONBody(w, r, &customerForm)

		if err != nil {
			log.Print("decoding product body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		customer, err := customerService.GetCustomerByCPF(r.Context(), customerForm.CPF)

		if err != nil {
			log.Print("get customer by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, customer)
	}
}

package handler

import (
	"log"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/domain/usecases"
	"github.com/thiagoluis88git/tech1/pkg/httpserver"
)

// @Summary Login
// @Description Login the customer by its CPF
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body dto.CustomerForm true "customer form"
// @Success 200 {object} dto.Token
// @Failure 404 "Customer not found"
// @Router /auth/login [post]
func LoginCustomerHandler(loginCustomerUseCase usecases.LoginCustomerUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customerForm dto.CustomerForm

		err := httpserver.DecodeJSONBody(w, r, &customerForm)

		if err != nil {
			log.Print("decoding customer form body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		token, err := loginCustomerUseCase.Execute(r.Context(), customerForm.CPF)

		if err != nil {
			log.Print("login user", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, token)
	}
}

// @Summary Login with unknown user
// @Description Login with unknown user. This is important if the user doesn't want to create an account
// @Tags Customer
// @Accept json
// @Produce json
// @Success 200 {object} dto.Token
// @Failure 404 "Customer not found"
// @Router /auth/login/unknown [post]
func LoginUnknownCustomerHandler(loginCustomerUseCase usecases.LoginUnknownCustomerUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := loginCustomerUseCase.Execute(r.Context())

		if err != nil {
			log.Print("login user", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, token)
	}
}

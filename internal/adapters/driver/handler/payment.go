package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/services"
	"github.com/thiagoluis88git/tech1/pkg/httpserver"
)

// @Summary Create new payment
// @Description Create a payment and return its ID. With it, we can proceed with a Order Creation
// @Tags Payment
// @Accept json
// @Produce json
// @Param product body domain.Payment true "payment"
// @Success 200 {object} domain.PaymentResponse
// @Failure 400 "Payment has required fields"
// @Router /api/payments [post]
func CreatePaymentHandler(productService *services.PaymentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var combo domain.Payment

		err := httpserver.DecodeJSONBody(w, r, &combo)

		if err != nil {
			log.Print("decoding payment body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		paymentResponse, err := productService.PayOrder(context.Background(), combo)

		if err != nil {
			log.Print("create payment", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, paymentResponse)
	}
}

// @Summary Get payment types
// @Description Get payment type, like [DEBIT, CREDIT]
// @Tags Payment
// @Accept json
// @Produce json
// @Success 200 {object} []string
// @Router /api/payments/type [get]
func GetPaymentTypeHandler(paymentService *services.PaymentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, paymentService.GetPaymentTypes())
	}
}

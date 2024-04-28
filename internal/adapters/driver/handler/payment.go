package handler

import (
	"context"
	"log"
	"net/http"
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/services"
	"thiagoluis88git/tech1/pkg/httpserver"
)

// @Summary Create new combo
// @Description Create new combo of products
// @Tags payment
// @Accept json
// @Produce json
// @Param product body domain.ComboForm true "combo"
// @Success 200 {object} domain.Payment
// @Failure 400 "Payment has required fields"
// @Router /api/payment [post]
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

func GetPaymentTypeHandler(paymentService *services.PaymentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, paymentService.GetPaymentTypes())
	}
}

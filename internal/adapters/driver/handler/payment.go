package handler

import (
	"net/http"
	"thiagoluis88git/tech1/internal/core/services"
	"thiagoluis88git/tech1/pkg/httpserver"
)

func GetPaymentTypeHandler(paymentService *services.PaymentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, paymentService.GetPaymentTypes())
	}
}

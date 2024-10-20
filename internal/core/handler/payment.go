package handler

import (
	"log"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/domain/usecases"
	"github.com/thiagoluis88git/tech1/pkg/httpserver"
)

// @Summary Create new payment
// @Description Create a payment and return its ID. With it, we can proceed with a Order Creation
// @Tags Payment
// @Accept json
// @Produce json
// @Param product body dto.Payment true "payment"
// @Success 200 {object} dto.PaymentResponse
// @Failure 400 "Payment has required fields"
// @Router /api/payments [post]
func CreatePaymentHandler(payOrder usecases.PayOrderUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payment dto.Payment

		err := httpserver.DecodeJSONBody(w, r, &payment)

		if err != nil {
			log.Print("decoding payment body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		paymentResponse, err := payOrder.Execute(r.Context(), payment)

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
// @Description Get payment type, like [DEBIT, CREDIT, QR Code (Mercado Pago)]
// @Tags Payment
// @Accept json
// @Produce json
// @Success 200 {object} []string
// @Router /api/payments/type [get]
func GetPaymentTypeHandler(getPaymentTypes usecases.GetPaymentTypesUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, getPaymentTypes.Execute())
	}
}

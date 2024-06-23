package handler

import (
	"log"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/core/domain"
	"github.com/thiagoluis88git/tech1/internal/core/services"
	"github.com/thiagoluis88git/tech1/pkg/environment"
	"github.com/thiagoluis88git/tech1/pkg/httpserver"
)

// @Summary Generate a QR Code
// @Description Generate a QR Code. This can be used to get the QR Code data, transform in a image and
// @Description pay with a Mercado Livre test account to activate a Webhook to proccess the order.
// @Tags QRCode
// @Accept json
// @Produce json
// @Param qrCodeForm body domain.QRCodeForm true "qrCodeForm"
// @Success 200 {object} domain.QRCodeDataResponse
// @Router /api/qrcode/generate [post]
func GenerateQRCodeHandler(qrCodeGeneratorService *services.QRCodeGeneratorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var form domain.Order

		err := httpserver.DecodeJSONBody(w, r, &form)

		if err != nil {
			log.Print("decoding qrcode body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		token := environment.GetQRCodeGatewayToken()
		response, err := qrCodeGeneratorService.GenerateQRCode(r.Context(), token, form)

		if err != nil {
			log.Print("generate qrcode", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, response)
	}
}

package dto

import "time"

type PaymentGatewayResponse struct {
	PaymentGatewayId string
	PaymentDate      time.Time
}

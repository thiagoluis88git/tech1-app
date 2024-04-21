package domain

import "time"

type PaymentGatewayResponse struct {
	PaymentGatewayId string
	PaymentDate      time.Time
}

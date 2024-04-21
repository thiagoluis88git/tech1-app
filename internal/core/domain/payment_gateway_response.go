package domain

import "time"

type PaymentGatewayResponse struct {
	PaymentGatewayId string    `json:"paymentGatewayId"`
	PaymentDate      time.Time `json:"paymentDate"`
}

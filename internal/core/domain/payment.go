package domain

import "time"

type Payment struct {
	CustomerID  *uint  `json:"customerId"`
	TotalPrice  int    `json:"totalPrice" validate:"required"`
	PaymentType string `json:"paymentType" validate:"required"`
}

type PaymentResponse struct {
	PaymentId        uint      `json:"paymentId"`
	PaymentGatewayId string    `json:"paymentGatewayId"`
	PaymentDate      time.Time `json:"paymentDate"`
}

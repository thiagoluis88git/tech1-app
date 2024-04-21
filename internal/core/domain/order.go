package domain

import "time"

type Order struct {
	OrderStatus  string
	TotalPrice   int            `json:"totalPrice" validate:"required"`
	CustomerID   *uint          `json:"customerId"`
	OrderProduct []OrderProduct `json:"orderProducts" validate:"required"`
	PaymentKind  string         `json:"paymentKind" validate:"required"`
	TickerId     int
}

type OrderProduct struct {
	ProductID uint `json:"productId" validate:"required"`
}

type OrderResponse struct {
	OrderId          uint
	OrderDate        time.Time `json:"productId"`
	TickerId         int       `json:"tickerId"`
	PaymentGatewayId string    `json:"paymentGatewayId"`
	CustomerName     *string   `json:"customerName"`
}

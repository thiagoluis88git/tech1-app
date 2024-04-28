package domain

import "time"

type Order struct {
	OrderStatus  string
	TotalPrice   int            `json:"totalPrice" validate:"required"`
	CustomerID   *uint          `json:"customerId"`
	PaymentID    uint           `json:"paymentId" validate:"required"`
	OrderProduct []OrderProduct `json:"orderProducts" validate:"required"`
	PaymentKind  string         `json:"paymentKind" validate:"required"`
	TickerId     int
}

type OrderProduct struct {
	ProductID uint `json:"productId" validate:"required"`
}

type OrderResponse struct {
	OrderId      uint      `json:"orderId"`
	OrderDate    time.Time `json:"orderDate"`
	TickerId     int       `json:"tickerId"`
	CustomerName *string   `json:"customerName"`
}

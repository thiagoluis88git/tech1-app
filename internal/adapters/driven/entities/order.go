package entities

import (
	"time"

	"gorm.io/gorm"
)

const (
	OrderStatusCreated      = "Criado"
	OrderStatusPreparing    = "Preparando"
	OrderStatusDone         = "Finalizado"
	OrderStatusDelivered    = "Entregue"
	OrderStatusNotDelivered = "Não entregue"
)

type Order struct {
	gorm.Model
	OrderStatus    string
	TotalPrice     int
	PaymentID      uint
	CustomerID     *uint
	Customer       *Customer
	TicketNumber   int
	PreparingAt    *time.Time
	DoneAt         *time.Time
	DeliveredAt    *time.Time
	NotDeliveredAt *time.Time
	OrderProduct   []OrderProduct
}

type OrderProduct struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Product   Product
}

type OrderTicketNumber struct {
	Date         int64 `gorm:"index;unique"`
	TicketNumber int
}

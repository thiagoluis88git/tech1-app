package entities

import (
	"time"

	"gorm.io/gorm"
)

const (
	OrderStatusCreated      = "Criado"
	OrderStatusPayed        = "Pago"
	OrderStatusPreparing    = "Preparando"
	OrderStatusDone         = "Finalizado"
	OrderStatusDelivered    = "Entregue"
	OrderStatusNotDelivered = "Não entregue"
)

type Order struct {
	gorm.Model
	OrderStatus    string
	TotalPrice     int
	CustomerID     *uint
	Customer       Customer
	TickerID       int
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
}

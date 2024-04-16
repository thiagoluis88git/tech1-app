package entities

import "gorm.io/gorm"

const (
	PaymentPayingStatus = "Pagando"
	PaymentPayedStatus  = "Pago"
	PaymentErrorStatus  = "Erro"
)

type PaymentOutbox struct {
	gorm.Model
	CustomerID    *uint
	Customer      *Customer
	OrderID       uint
	Order         Order
	PaymentStatus string
	// Retorno do gateway sobre qual tipo de pagamento o cliente selecionou
	PaymentKind string
}

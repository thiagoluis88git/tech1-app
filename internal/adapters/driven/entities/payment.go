package entities

import "gorm.io/gorm"

const (
	PaymentPayingStatus = "Pagando"
	PaymentPayedStatus  = "Pago"
	PaymentErrorStatus  = "Erro"

	PaymentCreditCardType = "Crédito"
	PaymentDebitType      = "Débito"
	PaymentVoucherType    = "Voucher"
)

type Payment struct {
	gorm.Model
	CustomerID    *uint
	Customer      *Customer
	OrderID       uint
	TotalPrice    int
	Order         Order
	PaymentStatus string
	PaymentKind   string
}

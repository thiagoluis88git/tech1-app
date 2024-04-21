package entities

import "gorm.io/gorm"

const (
	PaymentPayingStatus = "Pagando"
	PaymentPayedStatus  = "Pago"
	PaymentErrorStatus  = "Erro"

	PaymentCreditCardKind = "Crédito"
	PaymentDebitKind      = "Débito"
	PaymentVoucherKind    = "Voucher"
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

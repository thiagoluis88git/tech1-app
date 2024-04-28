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
	TotalPrice    int
	PaymentStatus string
	PaymentKind   string
}

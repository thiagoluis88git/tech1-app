package entities

import "gorm.io/gorm"

const (
	PaymentPayingStatus = "Pagando"
	PaymentPayedStatus  = "Pago"
	PaymentErrorStatus  = "Erro"

	PaymentCreditType = "Crédito"
	PaymentQRCodeType = "QR Code (Mercado Pago)"
)

type Payment struct {
	gorm.Model
	CustomerID    *uint
	Customer      *Customer
	TotalPrice    int
	PaymentStatus string
	PaymentType   string
}

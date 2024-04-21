package ports

import "thiagoluis88git/tech1/internal/core/domain"

type PaymentGateway interface {
	Pay(paymentResonse domain.PaymentResponse, payment domain.Payment) (domain.PaymentGatewayResponse, error)
}

package external

import (
	"thiagoluis88git/tech1/internal/core/domain"
	"thiagoluis88git/tech1/internal/core/ports"
	"time"

	"github.com/google/uuid"
)

type PaymentGateway struct {
}

func NewPaymentGateway() ports.PaymentGateway {
	return &PaymentGateway{}
}

func (p *PaymentGateway) Pay(paymentResonse domain.PaymentResponse, payment domain.Payment) (domain.PaymentGatewayResponse, error) {
	id := uuid.New()

	time.Sleep(3 * time.Second)

	return domain.PaymentGatewayResponse{
		PaymentGatewayId: id.String(),
		PaymentDate:      time.Now(),
	}, nil
}

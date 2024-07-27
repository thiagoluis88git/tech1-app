package external

import (
	"time"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/domain/repository"

	"github.com/google/uuid"
)

type PaymentGateway struct {
}

func NewPaymentGateway() repository.PaymentGateway {
	return &PaymentGateway{}
}

func (p *PaymentGateway) Pay(paymentResonse dto.PaymentResponse, payment dto.Payment) (dto.PaymentGatewayResponse, error) {
	id := uuid.New()

	time.Sleep(3 * time.Second)

	return dto.PaymentGatewayResponse{
		PaymentGatewayId: id.String(),
		PaymentDate:      time.Now(),
	}, nil
}

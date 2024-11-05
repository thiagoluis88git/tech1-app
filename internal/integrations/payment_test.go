package integrations_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/integrations"
)

func mockPaymentResponse() dto.PaymentResponse {
	return dto.PaymentResponse{}
}

func mockPayment() dto.Payment {
	return dto.Payment{}
}

func TestPayment(t *testing.T) {
	t.Parallel()

	t.Run("got success when pay with simulator", func(t *testing.T) {
		t.Parallel()

		sut := integrations.NewPaymentGateway()

		response, err := sut.Pay(mockPaymentResponse(), mockPayment())

		assert.NoError(t, err)
		assert.NotEmpty(t, response.PaymentGatewayId)
	})
}

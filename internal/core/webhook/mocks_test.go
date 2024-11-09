package webhook_test

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
)

type MockFinishOrderForQRCodeUseCase struct {
	mock.Mock
}

func (mock *MockFinishOrderForQRCodeUseCase) Execute(ctx context.Context, token string, form dto.ExternalPaymentEvent) error {
	args := mock.Called(ctx, token, form)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

package repositories_test

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1/internal/integrations/model"
)

type MockMercadoLivreDataSource struct {
	mock.Mock
}

func (mock *MockMercadoLivreDataSource) Generate(ctx context.Context, token string, input model.QRCodeInput) (string, error) {
	args := mock.Called(ctx, token, input)
	err := args.Error(1)

	if err != nil {
		return "", err
	}

	return args.Get(0).(string), nil
}

func (mock *MockMercadoLivreDataSource) GetPaymentData(ctx context.Context, token string, endpoint string) (model.MercadoLivrePaymentResponse, error) {
	args := mock.Called(ctx, token, endpoint)
	err := args.Error(1)

	if err != nil {
		return model.MercadoLivrePaymentResponse{}, err
	}

	return args.Get(0).(model.MercadoLivrePaymentResponse), nil
}

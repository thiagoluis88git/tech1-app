package repositories_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/integrations/model"
	"github.com/thiagoluis88git/tech1/internal/integrations/repositories"
	"github.com/thiagoluis88git/tech1/pkg/environment"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func setup() {
	os.Setenv(environment.QRCodeGatewayRootURL, "ROOT_URL")
	os.Setenv(environment.DBHost, "HOST")
	os.Setenv(environment.DBPort, "1234")
	os.Setenv(environment.DBUser, "User")
	os.Setenv(environment.DBPassword, "Pass")
	os.Setenv(environment.DBName, "Name")
	os.Setenv(environment.CognitoClientID, "ClienId")
	os.Setenv(environment.CognitoGroupAdmin, "Admin")
	os.Setenv(environment.CognitoGroupUser, "CognitoUser")
	os.Setenv(environment.CognitoUserPoolID, "USerPool")
	os.Setenv(environment.WebhookMercadoLivrePaymentURL, "WEBHOOK")
	os.Setenv(environment.QRCodeGatewayToken, "TOKEN")
	os.Setenv(environment.Region, "Region")
}

func mockOrder() dto.Order {
	return dto.Order{
		TotalPrice:   123.45,
		TicketNumber: 12,
		OrderProduct: []dto.OrderProduct{
			{
				ProductID:    uint(1),
				ProductPrice: 12.33,
			},
		},
	}
}

func TestMercadoLivreRepository(t *testing.T) {
	t.Parallel()
	setup()

	t.Run("got success when generating QR code repository", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		mockRepo := new(MockMercadoLivreDataSource)

		sut := repositories.NewMercadoLivreRepository(mockRepo)

		ctx := context.TODO()

		mockRepo.On("Generate", ctx, "token", mock.Anything).Return("QR_CODE", nil)

		response, err := sut.Generate(ctx, "token", mockOrder(), 5)

		assert.NoError(t, err)
		assert.Equal(t, "QR_CODE", response.Data)
	})

	t.Run("got error on Generate Repo when generating QR code repository", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		mockRepo := new(MockMercadoLivreDataSource)

		sut := repositories.NewMercadoLivreRepository(mockRepo)

		ctx := context.TODO()

		mockRepo.On("Generate", ctx, "token", mock.Anything).Return("", &responses.NetworkError{
			Code: 422,
		})

		response, err := sut.Generate(ctx, "token", mockOrder(), 5)

		assert.Error(t, err)
		assert.Empty(t, response.Data)
	})

	t.Run("got success when getting QR code payment data repository", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		mockRepo := new(MockMercadoLivreDataSource)

		sut := repositories.NewMercadoLivreRepository(mockRepo)

		ctx := context.TODO()

		mockRepo.On("GetPaymentData", ctx, "token", "endpoint").Return(model.MercadoLivrePaymentResponse{
			Status: "SUCCESS",
		}, nil)

		response, err := sut.GetQRCodePaymentData(ctx, "token", "endpoint")

		assert.NoError(t, err)
		assert.Equal(t, "SUCCESS", response.Status)
	})

	t.Run("got error on GetPaymentData Repository when getting QR code payment data repository", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		mockRepo := new(MockMercadoLivreDataSource)

		sut := repositories.NewMercadoLivreRepository(mockRepo)

		ctx := context.TODO()

		mockRepo.On("GetPaymentData", ctx, "token", "endpoint").Return(
			model.MercadoLivrePaymentResponse{},
			&responses.NetworkError{
				Code: 500,
			},
		)

		response, err := sut.GetQRCodePaymentData(ctx, "token", "endpoint")

		assert.Error(t, err)
		assert.Empty(t, response)
	})
}

package remote

import (
	"context"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/integrations/model"
	"github.com/thiagoluis88git/tech1/pkg/environment"
	"github.com/thiagoluis88git/tech1/pkg/httpserver"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type MercadoLivreDataSource interface {
	Generate(ctx context.Context, token string, input model.QRCodeInput) (string, error)
	GetPaymentData(ctx context.Context, token string, endpoint string) (model.MercadoLivrePaymentResponse, error)
}

type MercadoLivreRemoteDataSource struct {
	client   *http.Client
	endpoint string
}

func NewMercadoLivreDataSource(client *http.Client) MercadoLivreDataSource {
	return &MercadoLivreRemoteDataSource{
		client:   client,
		endpoint: environment.GetQRCodeGatewayRootURL(),
	}
}

func (ds *MercadoLivreRemoteDataSource) Generate(ctx context.Context, token string, input model.QRCodeInput) (string, error) {
	body, err := input.GetJSONBody()

	if err != nil {
		return "", &responses.NetworkError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	}

	response, err := httpserver.DoRequest(
		ctx,
		ds.client,
		ds.endpoint,
		&token,
		body,
		http.MethodPost,
		model.QRCodeData{},
	)

	if err != nil {
		return "", err
	}

	return response.QRData, nil
}

func (ds *MercadoLivreRemoteDataSource) GetPaymentData(ctx context.Context, token string, endpoint string) (model.MercadoLivrePaymentResponse, error) {
	response, err := httpserver.DoRequest(
		ctx,
		ds.client,
		endpoint,
		&token,
		nil,
		http.MethodGet,
		model.MercadoLivrePaymentResponse{},
	)

	if err != nil {
		return model.MercadoLivrePaymentResponse{}, err
	}

	return response, nil
}

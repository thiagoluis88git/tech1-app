package remote

import (
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/adapters/driven/external/model"
	"github.com/thiagoluis88git/tech1/pkg/environment"
	"github.com/thiagoluis88git/tech1/pkg/httpserver"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type QRCodeGeneratorDataSource interface {
	Generate(token string, input model.QRCodeInput) (string, error)
}

type MercadoLivreQRCOdeGeneratorRemoteDataSource struct {
	client   *http.Client
	endpoint string
}

func NewMercadoLivreQRCOdeGeneratorDataSource(client *http.Client) QRCodeGeneratorDataSource {
	return &MercadoLivreQRCOdeGeneratorRemoteDataSource{
		client:   client,
		endpoint: environment.GetQRCodeGatewayRootURL(),
	}
}

func (generator *MercadoLivreQRCOdeGeneratorRemoteDataSource) Generate(token string, input model.QRCodeInput) (string, error) {
	body, err := input.GetJSONBody()

	if err != nil {
		return "", &responses.NetworkError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	}

	response, err := httpserver.DoRequest(
		generator.client,
		generator.endpoint,
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

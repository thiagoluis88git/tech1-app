package remote_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/integrations/model"
	"github.com/thiagoluis88git/tech1/internal/integrations/remote"
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
	os.Setenv(environment.QRCodeGatewayToken, "token")
	os.Setenv(environment.Region, "Region")
}

type MockRoundTripper struct {
	Response *http.Response
}

func (trip *MockRoundTripper) RoundTrip(_ *http.Request) (*http.Response, error) {
	return trip.Response, nil
}

func TestMercadoLivreRemote(t *testing.T) {
	t.Parallel()
	setup()

	t.Run("got success when generate qrcode remote", func(t *testing.T) {
		t.Parallel()

		environment.LoadEnvironmentVariables()

		recorder := httptest.NewRecorder()
		recorder.Header().Add("Content-Type", "application/json")
		recorder.WriteHeader(http.StatusOK)
		_, err := recorder.WriteString(MockQRCode)

		assert.NoError(t, err)

		resultExpected := recorder.Result()

		mockClient := &http.Client{
			Transport: &MockRoundTripper{
				Response: resultExpected,
			},
		}

		ds := remote.NewMercadoLivreDataSource(mockClient)

		response, err := ds.Generate(context.TODO(), "token", model.QRCodeInput{})

		assert.NoError(t, err)
		assert.Equal(t, "QR_CODE", response)
	})

	t.Run("got error on invalid json when generate qrcode remote", func(t *testing.T) {
		t.Parallel()

		environment.LoadEnvironmentVariables()

		recorder := httptest.NewRecorder()
		recorder.Header().Add("Content-Type", "application/json")
		recorder.WriteHeader(http.StatusOK)
		_, err := recorder.WriteString("asddd{{}")

		assert.NoError(t, err)

		resultExpected := recorder.Result()

		mockClient := &http.Client{
			Transport: &MockRoundTripper{
				Response: resultExpected,
			},
		}

		ds := remote.NewMercadoLivreDataSource(mockClient)

		response, err := ds.Generate(context.TODO(), "token", model.QRCodeInput{})

		assert.Error(t, err)
		assert.Empty(t, response)

		var netError *responses.NetworkError
		isNetError := errors.As(err, &netError)
		assert.Equal(t, true, isNetError)
		assert.Equal(t, http.StatusUnprocessableEntity, netError.Code)
	})
}

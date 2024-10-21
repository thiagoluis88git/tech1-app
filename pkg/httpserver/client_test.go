package httpserver_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/pkg/httpserver"
)

func TestHttpClient(t *testing.T) {
	t.Parallel()

	t.Run("got success when creating http client", func(t *testing.T) {
		t.Parallel()

		client := httpserver.NewHTTPClient()

		assert.NotEmpty(t, client)
	})
}

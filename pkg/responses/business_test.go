package responses_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func TestBusinessResponse(t *testing.T) {
	t.Parallel()

	t.Run("got BadRequest error with Network Error when calling GetResponseError", func(t *testing.T) {
		t.Parallel()

		err := &responses.NetworkError{
			Code:    http.StatusBadRequest,
			Message: "BadRequest",
		}

		businessError := responses.GetResponseError(err, "MOCK")

		assert.Equal(t, "Bad request trying to execute MOCK - BadRequest", businessError.Error())
	})

	t.Run("got Unauthorized error with Network Error when calling GetResponseError", func(t *testing.T) {
		t.Parallel()

		err := &responses.NetworkError{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		}

		businessError := responses.GetResponseError(err, "MOCK")

		assert.Equal(t, "Unauthorized error trying to execute MOCK - Unauthorized", businessError.Error())
	})

	t.Run("got Unauthorized error with Network Error when calling GetResponseError", func(t *testing.T) {
		t.Parallel()

		err := &responses.NetworkError{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		}

		businessError := responses.GetResponseError(err, "MOCK")

		assert.Equal(t, "Unauthorized error trying to execute MOCK - Unauthorized", businessError.Error())
	})

	t.Run("got Forbidden error with Network Error when calling GetResponseError", func(t *testing.T) {
		t.Parallel()

		err := &responses.NetworkError{
			Code:    http.StatusForbidden,
			Message: "Forbidden",
		}

		businessError := responses.GetResponseError(err, "MOCK")

		assert.Equal(t, "Forbiden error trying to execute MOCK - Forbidden", businessError.Error())
	})

	t.Run("got NotFound error with Network Error when calling GetResponseError", func(t *testing.T) {
		t.Parallel()

		err := &responses.NetworkError{
			Code:    http.StatusNotFound,
			Message: "NotFound",
		}

		businessError := responses.GetResponseError(err, "MOCK")

		assert.Equal(t, "Not found trying to execute MOCK - NotFound", businessError.Error())
	})

	t.Run("got Conflict error with Network Error when calling GetResponseError", func(t *testing.T) {
		t.Parallel()

		err := &responses.NetworkError{
			Code:    http.StatusConflict,
			Message: "Conflict",
		}

		businessError := responses.GetResponseError(err, "MOCK")

		assert.Equal(t, "Conflit with some data using the service MOCK - Conflict", businessError.Error())
	})

	t.Run("got UnprocessableEntity error with Network Error when calling GetResponseError", func(t *testing.T) {
		t.Parallel()

		err := &responses.NetworkError{
			Code:    http.StatusUnprocessableEntity,
			Message: "UnprocessableEntity",
		}

		businessError := responses.GetResponseError(err, "MOCK")

		assert.Equal(t, "Logic error found in service MOCK - UnprocessableEntity", businessError.Error())
	})

	t.Run("got InternalServerError error with Network Error when calling GetResponseError", func(t *testing.T) {
		t.Parallel()

		err := &responses.NetworkError{
			Code:    http.StatusInternalServerError,
			Message: "InternalServerError",
		}

		businessError := responses.GetResponseError(err, "MOCK")

		assert.Equal(t, "Unexpected internal error trying to execute service MOCK - InternalServerError", businessError.Error())
	})
}

package httpserver_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/pkg/httpserver"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func TestRequestResponseHelper(t *testing.T) {
	t.Parallel()

	t.Run("got error when passing empty json", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var destination any
			err := httpserver.DecodeJSONBody(w, r, &destination)

			assert.Error(t, err)
			assert.Equal(t, "Request body must not be empty", err.Error())

			w.WriteHeader(http.StatusOK)
		})

		ts := httptest.NewServer(responseHandler)
		defer ts.Close()

		req, _ := http.NewRequest(http.MethodPost, ts.URL+"/mock", strings.NewReader(""))
		req.Header.Add("Content-Type", "application/json")

		response, err := ts.Client().Do(req)

		assert.NoError(t, err)
		defer response.Body.Close()
	})

	t.Run("got error when not passing Content-Type", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var destination any
			err := httpserver.DecodeJSONBody(w, r, &destination)

			assert.Error(t, err)
			assert.Equal(t, "Content-Type header is not application/json", err.Error())

			w.WriteHeader(http.StatusOK)
		})

		ts := httptest.NewServer(responseHandler)
		defer ts.Close()

		req, _ := http.NewRequest(http.MethodPost, ts.URL+"/mock", strings.NewReader(""))

		response, err := ts.Client().Do(req)

		assert.NoError(t, err)
		defer response.Body.Close()
	})

	t.Run("got error when passing empty Content-Type", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var destination any
			err := httpserver.DecodeJSONBody(w, r, &destination)

			assert.Error(t, err)
			assert.Equal(t, "Content-Type header is not application/json", err.Error())

			w.WriteHeader(http.StatusOK)
		})

		ts := httptest.NewServer(responseHandler)
		defer ts.Close()

		req, _ := http.NewRequest(http.MethodPost, ts.URL+"/mock", strings.NewReader(""))
		req.Header.Add("Content-Type", "")

		response, err := ts.Client().Do(req)

		assert.NoError(t, err)
		defer response.Body.Close()
	})

	t.Run("got error when passing badly-formed json", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var destination any
			err := httpserver.DecodeJSONBody(w, r, &destination)

			assert.Error(t, err)
			assert.Equal(t, "Request body contains badly-formed JSON", err.Error())

			w.WriteHeader(http.StatusOK)
		})

		ts := httptest.NewServer(responseHandler)
		defer ts.Close()

		req, _ := http.NewRequest(http.MethodPost, ts.URL+"/mock", strings.NewReader(`{"teste": "ttt}`))
		req.Header.Add("Content-Type", "application/json")

		response, err := ts.Client().Do(req)

		assert.NoError(t, err)
		defer response.Body.Close()
	})

	t.Run("got error when passing badly-formed at position json", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var destination any
			err := httpserver.DecodeJSONBody(w, r, &destination)

			assert.Error(t, err)
			assert.Contains(t, err.Error(), "Request body contains badly-formed JSON (at position")

			w.WriteHeader(http.StatusOK)
		})

		ts := httptest.NewServer(responseHandler)
		defer ts.Close()

		req, _ := http.NewRequest(http.MethodPost, ts.URL+"/mock", strings.NewReader(`{"teste": "teste"{}}`))
		req.Header.Add("Content-Type", "application/json")

		response, err := ts.Client().Do(req)

		assert.NoError(t, err)
		defer response.Body.Close()
	})

	t.Run("got error when passing unknown fields", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var destination dto.Token
			err := httpserver.DecodeJSONBody(w, r, &destination)

			assert.Error(t, err)
			assert.Contains(t, err.Error(), "Request body contains unknown field")

			w.WriteHeader(http.StatusOK)
		})

		ts := httptest.NewServer(responseHandler)
		defer ts.Close()

		req, _ := http.NewRequest(http.MethodPost, ts.URL+"/mock", strings.NewReader(`{"teste": "teste"}`))
		req.Header.Add("Content-Type", "application/json")

		response, err := ts.Client().Do(req)

		assert.NoError(t, err)
		defer response.Body.Close()
	})

	t.Run("got error when passing data without required fields", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var destination dto.ComboForm
			err := httpserver.DecodeJSONBody(w, r, &destination)

			assert.Error(t, err)
			assert.Contains(t, err.Error(), "Error JSON required fields")

			w.WriteHeader(http.StatusOK)
		})

		ts := httptest.NewServer(responseHandler)
		defer ts.Close()

		req, _ := http.NewRequest(http.MethodPost, ts.URL+"/mock", strings.NewReader(`{"name": "Hamburguer"}`))
		req.Header.Add("Content-Type", "application/json")

		response, err := ts.Client().Do(req)

		assert.NoError(t, err)
		defer response.Body.Close()
	})

	t.Run("got success when passing valid path param", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := httpserver.GetPathParamFromRequest(r, "id")
			assert.NoError(t, err)
			assert.NotEmpty(t, id)
			assert.Equal(t, "123", id)

			w.WriteHeader(http.StatusOK)
		})

		router := chi.NewRouter()

		router.Get("/product/{id}", responseHandler)

		ts := httptest.NewServer(router)

		defer ts.Close()

		req, _ := http.NewRequest(http.MethodGet, ts.URL+"/product/123", nil)
		req.Header.Add("Content-Type", "application/json")

		response, err := ts.Client().Do(req)

		assert.NoError(t, err)
		defer response.Body.Close()
	})

	t.Run("got error when passing invalid path param", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := httpserver.GetPathParamFromRequest(r, "idinvalid")
			assert.Error(t, err)
			assert.Empty(t, id)

			var netError *responses.NetworkError
			isNestError := errors.As(err, &netError)

			assert.True(t, isNestError)
			assert.Equal(t, http.StatusUnprocessableEntity, netError.Code)
			assert.Equal(t, "idinvalid param not found in path", netError.Message)

			w.WriteHeader(http.StatusUnprocessableEntity)
		})

		router := chi.NewRouter()

		router.Get("/product/{id}", responseHandler)

		ts := httptest.NewServer(router)

		defer ts.Close()

		req, _ := http.NewRequest(http.MethodGet, ts.URL+"/product/123", nil)
		req.Header.Add("Content-Type", "application/json")

		response, err := ts.Client().Do(req)

		assert.NoError(t, err)
		defer response.Body.Close()
	})
}

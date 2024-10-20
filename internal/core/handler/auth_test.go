package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/handler"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func mockLoginCustomer() dto.CustomerForm {
	return dto.CustomerForm{
		CPF: "83212446293",
	}
}
func TestAuthHandler(t *testing.T) {
	t.Parallel()

	t.Run("got success when calling login handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockLoginCustomer())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/auth/login", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		loginCustomerUseCase := new(MockLoginCustomerUseCase)

		loginCustomerUseCase.On("Execute", req.Context(), "83212446293").Return(dto.Token{
			AccessToken: "eYmly",
		}, nil)

		loginCustomerHandler := handler.LoginCustomerHandler(loginCustomerUseCase)

		loginCustomerHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var customer dto.Token
		err = json.Unmarshal(recorder.Body.Bytes(), &customer)

		assert.NoError(t, err)

		assert.Equal(t, "eYmly", customer.AccessToken)
	})

	t.Run("got error on UseCase when calling login handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockLoginCustomer())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/auth/login", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		loginCustomerUseCase := new(MockLoginCustomerUseCase)

		loginCustomerUseCase.On("Execute", req.Context(), "83212446293").Return(dto.Token{}, &responses.BusinessResponse{
			StatusCode: 500,
		})

		loginCustomerHandler := handler.LoginCustomerHandler(loginCustomerUseCase)

		loginCustomerHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("got error with invalid json when calling login handler", func(t *testing.T) {
		t.Parallel()

		body := bytes.NewBuffer([]byte("asdr{{}"))

		req := httptest.NewRequest(http.MethodPost, "/auth/login", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		loginCustomerUseCase := new(MockLoginCustomerUseCase)

		loginCustomerUseCase.On("Execute", req.Context(), "83212446293").Return(dto.Token{
			AccessToken: "eYmly",
		}, nil)

		loginCustomerHandler := handler.LoginCustomerHandler(loginCustomerUseCase)

		loginCustomerHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

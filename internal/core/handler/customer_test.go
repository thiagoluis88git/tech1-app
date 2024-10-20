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

func mockCustomer() map[string]any {
	return map[string]any{
		"name":  "Teste",
		"cpf":   "83212446293",
		"email": "teste@gmail.com",
	}
}

func TestCustomerHandler(t *testing.T) {
	t.Parallel()

	t.Run("got success when calling create customer handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCustomer())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/customer", body)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer eyAfgg")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		createCustomerUseCase := new(MockCreateCustomerUseCase)

		createCustomerUseCase.On("Execute", req.Context(), dto.Customer{
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}).Return(dto.CustomerResponse{
			Id: uint(1),
		}, nil)

		createCustomerHandler := handler.CreateCustomerHandler(createCustomerUseCase)

		createCustomerHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("got error on UseCase when calling create customer handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCustomer())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/customer", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		createCustomerUseCase := new(MockCreateCustomerUseCase)

		createCustomerUseCase.On("Execute", req.Context(), dto.Customer{
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}).Return(dto.CustomerResponse{}, &responses.BusinessResponse{
			StatusCode: 500,
		})

		createCustomerHandler := handler.CreateCustomerHandler(createCustomerUseCase)

		createCustomerHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("got error with invalid json data when calling create customer handler", func(t *testing.T) {
		t.Parallel()

		body := bytes.NewBuffer([]byte("sss{{}"))

		req := httptest.NewRequest(http.MethodPost, "/api/customer", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		createCustomerUseCase := new(MockCreateCustomerUseCase)

		createCustomerUseCase.On("Execute", req.Context(), dto.Customer{
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}).Return(dto.CustomerResponse{
			Id: uint(1),
		}, nil)

		createCustomerHandler := handler.CreateCustomerHandler(createCustomerUseCase)

		createCustomerHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

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

func mockCustomerByCPF() map[string]any {
	return map[string]any{
		"cpf": "83212446293",
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

	t.Run("got success when calling update customer handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCustomer())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/customer/{id}", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "123")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		updateCustomerUseCase := new(MockUpdateCustomerUseCase)

		updateCustomerUseCase.On("Execute", req.Context(), dto.Customer{
			ID:    uint(123),
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}).Return(nil)

		updateCustomerHandler := handler.UpdateCustomerHandler(updateCustomerUseCase)

		updateCustomerHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})

	t.Run("got error without param when calling update customer handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCustomer())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/customer/{id}", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		updateCustomerUseCase := new(MockUpdateCustomerUseCase)

		updateCustomerUseCase.On("Execute", req.Context(), dto.Customer{
			ID:    uint(123),
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}).Return(nil)

		updateCustomerHandler := handler.UpdateCustomerHandler(updateCustomerUseCase)

		updateCustomerHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("got error with invalid param when calling update customer handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCustomer())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/customer/{id}", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "123xxc")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		updateCustomerUseCase := new(MockUpdateCustomerUseCase)

		updateCustomerUseCase.On("Execute", req.Context(), dto.Customer{
			ID:    uint(123),
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}).Return(nil)

		updateCustomerHandler := handler.UpdateCustomerHandler(updateCustomerUseCase)

		updateCustomerHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("got error calling UseCase when calling update customer handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCustomer())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/customer/{id}", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "123")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		updateCustomerUseCase := new(MockUpdateCustomerUseCase)

		updateCustomerUseCase.On("Execute", req.Context(), dto.Customer{
			ID:    uint(123),
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}).Return(&responses.BusinessResponse{
			StatusCode: 500,
		})

		updateCustomerHandler := handler.UpdateCustomerHandler(updateCustomerUseCase)

		updateCustomerHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("got error with invalid json when calling update customer handler", func(t *testing.T) {
		t.Parallel()

		body := bytes.NewBuffer([]byte("ffdd{{}"))

		req := httptest.NewRequest(http.MethodPost, "/api/customer/{id}", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "123xxc")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		updateCustomerUseCase := new(MockUpdateCustomerUseCase)

		updateCustomerUseCase.On("Execute", req.Context(), dto.Customer{
			ID:    uint(123),
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}).Return(nil)

		updateCustomerHandler := handler.UpdateCustomerHandler(updateCustomerUseCase)

		updateCustomerHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("got success when calling get customer by id handler", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodGet, "/api/customer/{id}", nil)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "123")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getCustomerById := new(MockGetCustomerByIdUseCase)

		getCustomerById.On("Execute", req.Context(), uint(123)).Return(dto.Customer{
			ID:    uint(123),
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}, nil)

		getCustomerByIdHandler := handler.GetCustomerByIdHandler(getCustomerById)

		getCustomerByIdHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var customer dto.Customer
		err := json.Unmarshal(recorder.Body.Bytes(), &customer)

		assert.NoError(t, err)

		assert.Equal(t, "Teste", customer.Name)
	})

	t.Run("got error without param when calling get customer by id handler", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodGet, "/api/customer/{id}", nil)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getCustomerById := new(MockGetCustomerByIdUseCase)

		getCustomerById.On("Execute", req.Context(), uint(123)).Return(dto.Customer{
			ID:    uint(123),
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}, nil)

		getCustomerByIdHandler := handler.GetCustomerByIdHandler(getCustomerById)

		getCustomerByIdHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("got error with invalid param when calling get customer by id handler", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodGet, "/api/customer/{id}", nil)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "123srvb")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getCustomerById := new(MockGetCustomerByIdUseCase)

		getCustomerById.On("Execute", req.Context(), uint(123)).Return(dto.Customer{
			ID:    uint(123),
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}, nil)

		getCustomerByIdHandler := handler.GetCustomerByIdHandler(getCustomerById)

		getCustomerByIdHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("got error with UseCase when calling get customer by id handler", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodGet, "/api/customer/{id}", nil)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "123")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getCustomerById := new(MockGetCustomerByIdUseCase)

		getCustomerById.On("Execute", req.Context(), uint(123)).Return(dto.Customer{}, &responses.BusinessResponse{
			StatusCode: 500,
		})

		getCustomerByIdHandler := handler.GetCustomerByIdHandler(getCustomerById)

		getCustomerByIdHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("got success when calling get customer by cpf handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCustomerByCPF())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/customer/login", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getCustomerByCPF := new(MockGetCustomerByCPFUseCase)

		getCustomerByCPF.On("Execute", req.Context(), "83212446293").Return(dto.Customer{
			ID:    uint(123),
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}, nil)

		getCustomerCPFIdHandler := handler.GetCustomerByCPFHandler(getCustomerByCPF)

		getCustomerCPFIdHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var customer dto.Customer
		err = json.Unmarshal(recorder.Body.Bytes(), &customer)

		assert.NoError(t, err)

		assert.Equal(t, "Teste", customer.Name)
	})

	t.Run("got error on UseCase when calling get customer by cpf handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCustomerByCPF())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/customer/login", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getCustomerByCPF := new(MockGetCustomerByCPFUseCase)

		getCustomerByCPF.On("Execute", req.Context(), "83212446293").Return(dto.Customer{}, &responses.BusinessResponse{
			StatusCode: 500,
		})

		getCustomerCPFIdHandler := handler.GetCustomerByCPFHandler(getCustomerByCPF)

		getCustomerCPFIdHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("got error with invalid json when calling get customer by cpf handler", func(t *testing.T) {
		t.Parallel()

		body := bytes.NewBuffer([]byte("assd{{}"))

		req := httptest.NewRequest(http.MethodPost, "/api/customer/login", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getCustomerByCPF := new(MockGetCustomerByCPFUseCase)

		getCustomerByCPF.On("Execute", req.Context(), "83012446293").Return(dto.Customer{
			ID:    uint(123),
			Name:  "Teste",
			CPF:   "83212446293",
			Email: "teste@gmail.com",
		}, nil)

		getCustomerCPFIdHandler := handler.GetCustomerByCPFHandler(getCustomerByCPF)

		getCustomerCPFIdHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

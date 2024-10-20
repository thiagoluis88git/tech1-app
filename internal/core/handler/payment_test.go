package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/handler"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func mockPayment() dto.Payment {
	return dto.Payment{
		TotalPrice:  123.32,
		PaymentType: "CREDIT",
	}
}
func TestPaymentHandler(t *testing.T) {
	t.Parallel()

	t.Run("got success when calling create payment handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockPayment())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/payment", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		payUseCase := new(MockPayOrderUseCase)

		payUseCase.On("Execute", req.Context(), dto.Payment{
			TotalPrice:  123.32,
			PaymentType: "CREDIT",
		}).Return(dto.PaymentResponse{
			PaymentId:        uint(12),
			PaymentGatewayId: "gatewayID",
			PaymentDate:      time.Date(2024, time.March, 12, 20, 20, 0, 0, time.Local),
		}, nil)

		createPaymentHandler := handler.CreatePaymentHandler(payUseCase)

		createPaymentHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var paymentResponse dto.PaymentResponse
		err = json.Unmarshal(recorder.Body.Bytes(), &paymentResponse)

		assert.NoError(t, err)

		assert.Equal(t, uint(12), paymentResponse.PaymentId)
		assert.Equal(t, "gatewayID", paymentResponse.PaymentGatewayId)
	})

	t.Run("got error on UseCase when calling create payment handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockPayment())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/payment", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		payUseCase := new(MockPayOrderUseCase)

		payUseCase.On("Execute", req.Context(), dto.Payment{
			TotalPrice:  123.32,
			PaymentType: "CREDIT",
		}).Return(dto.PaymentResponse{}, &responses.BusinessResponse{
			StatusCode: 500,
		})

		createPaymentHandler := handler.CreatePaymentHandler(payUseCase)

		createPaymentHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("got error with invalid json when calling create payment handler", func(t *testing.T) {
		t.Parallel()

		body := bytes.NewBuffer([]byte("sdrt{{}"))

		req := httptest.NewRequest(http.MethodPost, "/api/payment", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		payUseCase := new(MockPayOrderUseCase)

		payUseCase.On("Execute", req.Context(), dto.Payment{
			TotalPrice:  123.32,
			PaymentType: "CREDIT",
		}).Return(dto.PaymentResponse{}, &responses.BusinessResponse{
			StatusCode: 500,
		})

		createPaymentHandler := handler.CreatePaymentHandler(payUseCase)

		createPaymentHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("got success when calling get payment types handler", func(t *testing.T) {
		t.Parallel()

		body := bytes.NewBuffer([]byte("sdrt{{}"))

		req := httptest.NewRequest(http.MethodPost, "/api/payment/types", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getPaymentTypesUseCase := new(MockGetPaymentTypesUseCase)

		getPaymentTypesUseCase.On("Execute").Return([]string{
			"CREDIT",
			"MERCADO PAGO",
		}, nil)

		createPaymentHandler := handler.GetPaymentTypeHandler(getPaymentTypesUseCase)

		createPaymentHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var types []string
		err := json.Unmarshal(recorder.Body.Bytes(), &types)

		assert.NoError(t, err)

		assert.Equal(t, "CREDIT", types[0])
		assert.Equal(t, "MERCADO PAGO", types[1])
	})
}

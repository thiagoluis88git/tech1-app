package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/handler"
	"github.com/thiagoluis88git/tech1/pkg/environment"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func mockQRCodeOrder() dto.QRCodeOrder {
	return dto.QRCodeOrder{
		TotalPrice:   123.45,
		TicketNumber: 12,
		OrderProduct: []dto.OrderProduct{
			{
				ProductID:    uint(1),
				ProductPrice: 34.56,
			},
		},
	}
}

func setup() {
	os.Setenv(environment.QRCodeGatewayToken, "token")
}

func TestGenerateQRCodeHandler(t *testing.T) {
	t.Parallel()
	setup()

	t.Run("got success when calling generate qrcode handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockQRCodeOrder())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/qrcode", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		generateQRCodeUseCase := new(MockGenerateQRCodePaymentUseCase)

		now := time.Now()
		orderDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

		var wg sync.WaitGroup
		channel := make(chan bool, 1)
		wg.Add(1)

		generateQRCodeUseCase.On("Execute", req.Context(), "token", mockQRCodeOrder(), orderDate.UnixMilli(), &wg, channel).
			Return(dto.QRCodeDataResponse{
				Data: "QR_CODE",
			}, nil)

		generateQRCodeHandler := handler.GenerateQRCodeHandler(generateQRCodeUseCase)

		generateQRCodeHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response dto.QRCodeDataResponse
		err = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.NoError(t, err)

		assert.Equal(t, "QR_CODE", response.Data)
	})

	t.Run("got error on GenerateQRCode UseCase when calling generate qrcode handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockQRCodeOrder())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/qrcode", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		generateQRCodeUseCase := new(MockGenerateQRCodePaymentUseCase)

		now := time.Now()
		orderDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

		var wg sync.WaitGroup
		channel := make(chan bool, 1)
		wg.Add(1)

		generateQRCodeUseCase.On("Execute", req.Context(), "token", mockQRCodeOrder(), orderDate.UnixMilli(), &wg, channel).
			Return(dto.QRCodeDataResponse{}, &responses.BusinessResponse{
				StatusCode: 500,
			})

		generateQRCodeHandler := handler.GenerateQRCodeHandler(generateQRCodeUseCase)

		generateQRCodeHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("got error on invalid json UseCase when calling generate qrcode handler", func(t *testing.T) {
		t.Parallel()

		body := bytes.NewBuffer([]byte("sdfg{{}"))

		req := httptest.NewRequest(http.MethodPost, "/api/qrcode", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		generateQRCodeUseCase := new(MockGenerateQRCodePaymentUseCase)

		generateQRCodeHandler := handler.GenerateQRCodeHandler(generateQRCodeUseCase)

		generateQRCodeHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

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

func mockCreateUserForm() dto.UserAdmin {
	return dto.UserAdmin{
		Name:  "Name",
		CPF:   "12345678910",
		Email: "teste@email.com",
	}
}

func TestUserAdminHandler(t *testing.T) {
	t.Parallel()
	setup()

	t.Run("got success when calling create user admin handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCreateUserForm())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/auth/user", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		createUserUseCase := new(MockCreateUserUseCase)

		createUserUseCase.On("Execute", req.Context(), mockCreateUserForm()).
			Return(dto.UserAdminResponse{
				Id: uint(2),
			}, nil)

		createUserHandler := handler.CreateUserHandler(createUserUseCase)

		createUserHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response dto.UserAdminResponse
		err = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.NoError(t, err)

		assert.Equal(t, uint(2), response.Id)
	})

	t.Run("got error on CreateUser UseCase when calling create user admin handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCreateUserForm())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/auth/user", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		createUserUseCase := new(MockCreateUserUseCase)

		createUserUseCase.On("Execute", req.Context(), mockCreateUserForm()).
			Return(dto.UserAdminResponse{}, &responses.BusinessResponse{
				StatusCode: 503,
			})

		createUserHandler := handler.CreateUserHandler(createUserUseCase)

		createUserHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusServiceUnavailable, recorder.Code)
	})

	t.Run("got error on invalid json when calling create user admin handler", func(t *testing.T) {
		t.Parallel()

		body := bytes.NewBuffer([]byte("afff{{}"))

		req := httptest.NewRequest(http.MethodPost, "/auth/user", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		createUserUseCase := new(MockCreateUserUseCase)

		createUserHandler := handler.CreateUserHandler(createUserUseCase)

		createUserHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

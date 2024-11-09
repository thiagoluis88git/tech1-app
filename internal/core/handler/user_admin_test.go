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

	t.Run("got success when calling update user admin handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCreateUserForm())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPut, "/auth/user/{id}", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "3")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		updateUserUseCase := new(MockUpdateUserUseCase)

		updateUserUseCase.On("Execute", req.Context(), mockUpdateUserForm()).
			Return(nil)

		createUserHandler := handler.UpdateUserHandler(updateUserUseCase)

		createUserHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})

	t.Run("got error on UpdateUser UseCase when calling update user admin handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCreateUserForm())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPut, "/auth/user/{id}", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "3")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		updateUserUseCase := new(MockUpdateUserUseCase)

		updateUserUseCase.On("Execute", req.Context(), mockUpdateUserForm()).
			Return(&responses.BusinessResponse{
				StatusCode: 409,
			})

		createUserHandler := handler.UpdateUserHandler(updateUserUseCase)

		createUserHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusConflict, recorder.Code)
	})

	t.Run("got error on invalid json UseCase when calling update user admin handler", func(t *testing.T) {
		t.Parallel()

		body := bytes.NewBuffer([]byte("sdfff{{}"))

		req := httptest.NewRequest(http.MethodPut, "/auth/user/{id}", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "3")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		updateUserUseCase := new(MockUpdateUserUseCase)

		createUserHandler := handler.UpdateUserHandler(updateUserUseCase)

		createUserHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("got error on invalid id when calling update user admin handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCreateUserForm())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPut, "/auth/user/{id}", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "x3")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		updateUserUseCase := new(MockUpdateUserUseCase)

		createUserHandler := handler.UpdateUserHandler(updateUserUseCase)

		createUserHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("got error on missing id when calling update user admin handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockCreateUserForm())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPut, "/auth/user/{id}", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		updateUserUseCase := new(MockUpdateUserUseCase)

		createUserHandler := handler.UpdateUserHandler(updateUserUseCase)

		createUserHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("got success when calling get user by id user admin handler", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodGet, "/auth/user/{id}", nil)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "3")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getUserByIDUseCase := new(MockGetUserByIdUseCase)

		getUserByIDUseCase.On("Execute", req.Context(), uint(3)).
			Return(dto.UserAdmin{
				ID:   uint(3),
				Name: "Test Name",
			}, nil)

		createUserHandler := handler.GetUserByIdHandler(getUserByIDUseCase)

		createUserHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response dto.UserAdmin
		err := json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.NoError(t, err)

		assert.Equal(t, uint(3), response.ID)
		assert.Equal(t, "Test Name", response.Name)
	})

	t.Run("got error on GetUser UseCase when calling get user by id user admin handler", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodGet, "/auth/user/{id}", nil)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "3")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getUserByIDUseCase := new(MockGetUserByIdUseCase)

		getUserByIDUseCase.On("Execute", req.Context(), uint(3)).
			Return(dto.UserAdmin{}, &responses.BusinessResponse{
				StatusCode: 404,
			})

		createUserHandler := handler.GetUserByIdHandler(getUserByIDUseCase)

		createUserHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("got error on invalid id UseCase when calling get user by id user admin handler", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodGet, "/auth/user/{id}", nil)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "x3")

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getUserByIDUseCase := new(MockGetUserByIdUseCase)

		createUserHandler := handler.GetUserByIdHandler(getUserByIDUseCase)

		createUserHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("got error on missing id UseCase when calling get user by id user admin handler", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodGet, "/auth/user/{id}", nil)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getUserByIDUseCase := new(MockGetUserByIdUseCase)

		createUserHandler := handler.GetUserByIdHandler(getUserByIDUseCase)

		createUserHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("got success when calling get user by CPF user admin handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockGetUserByCPF())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/users/login", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getUserByCPFUseCase := new(MockGetUserByCPFUseCase)

		getUserByCPFUseCase.On("Execute", req.Context(), "12345678910").
			Return(dto.UserAdmin{
				ID:   uint(3),
				Name: "Test Name",
			}, nil)

		handler := handler.GetUserByCPFHandler(getUserByCPFUseCase)

		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response dto.UserAdmin
		err = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.NoError(t, err)

		assert.Equal(t, uint(3), response.ID)
		assert.Equal(t, "Test Name", response.Name)
	})

	t.Run("got error on GetUser UseCase when calling get user by CPF user admin handler", func(t *testing.T) {
		t.Parallel()

		jsonData, err := json.Marshal(mockGetUserByCPF())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/users/login", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getUserByCPFUseCase := new(MockGetUserByCPFUseCase)

		getUserByCPFUseCase.On("Execute", req.Context(), "12345678910").
			Return(dto.UserAdmin{}, &responses.BusinessResponse{
				StatusCode: 404,
			})

		handler := handler.GetUserByCPFHandler(getUserByCPFUseCase)

		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("got error on invalid json when calling get user by CPF user admin handler", func(t *testing.T) {
		t.Parallel()

		body := bytes.NewBuffer([]byte("dff{{}"))

		req := httptest.NewRequest(http.MethodPost, "/users/login", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		getUserByCPFUseCase := new(MockGetUserByCPFUseCase)

		handler := handler.GetUserByCPFHandler(getUserByCPFUseCase)

		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

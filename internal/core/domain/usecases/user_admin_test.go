package usecases

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

func mockUserAdmin() dto.UserAdmin {
	return dto.UserAdmin{
		CPF: "83212446293",
	}
}

func newUserAdmin() dto.UserAdmin {
	return dto.UserAdmin{
		CPF: "832.124.462-93",
	}
}

func newInvalidUserAdmin() dto.UserAdmin {
	return dto.UserAdmin{
		CPF: "830.124.462-93",
	}
}

func TestUserAdminUseCases(t *testing.T) {
	t.Parallel()

	t.Run("got success when creating user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewCreateUserUseCase(NewValidateCPFUseCase(), mockUserAdminRepository)

		ctx := context.TODO()

		mockUserAdminRepository.On("CreateUser", ctx, mockUserAdmin()).Return(uint(2), nil)

		response, err := sut.Execute(ctx, newUserAdmin())

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		assert.Equal(t, uint(2), response.Id)
	})

	t.Run("got error on Create Use Repo when creating user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewCreateUserUseCase(NewValidateCPFUseCase(), mockUserAdminRepository)

		ctx := context.TODO()

		mockUserAdminRepository.On("CreateUser", ctx, mockUserAdmin()).Return(uint(0), &responses.NetworkError{
			Code: 400,
		})

		response, err := sut.Execute(ctx, newUserAdmin())

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got error on Validate CPF when creating user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewCreateUserUseCase(NewValidateCPFUseCase(), mockUserAdminRepository)

		ctx := context.TODO()

		response, err := sut.Execute(ctx, newInvalidUserAdmin())

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got success when getting user by cpf user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewGetUserByCPFUseCase(NewValidateCPFUseCase(), mockUserAdminRepository)

		ctx := context.TODO()

		mockUserAdminRepository.On("GetUserByCPF", ctx, "83212446293").Return(dto.UserAdmin{
			ID:   uint(2),
			Name: "User Name",
		}, nil)

		response, err := sut.Execute(ctx, "832.124.462-93")

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		assert.Equal(t, uint(2), response.ID)
		assert.Equal(t, "User Name", response.Name)
	})

	t.Run("got error on GetUserByCPF Repository when getting user by cpf user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewGetUserByCPFUseCase(NewValidateCPFUseCase(), mockUserAdminRepository)

		ctx := context.TODO()

		mockUserAdminRepository.On("GetUserByCPF", ctx, "83212446293").
			Return(dto.UserAdmin{}, &responses.NetworkError{
				Code: 401,
			})

		response, err := sut.Execute(ctx, "832.124.462-93")

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got error on validate cpf when getting user by cpf user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewGetUserByCPFUseCase(NewValidateCPFUseCase(), mockUserAdminRepository)

		ctx := context.TODO()

		response, err := sut.Execute(ctx, "830.124.462-93")

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got success when getting user by id user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewGetUserByIdUseCase(mockUserAdminRepository)

		ctx := context.TODO()

		mockUserAdminRepository.On("GetUserById", ctx, uint(2)).Return(dto.UserAdmin{
			ID:   uint(2),
			Name: "User Name",
		}, nil)

		response, err := sut.Execute(ctx, uint(2))

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		assert.Equal(t, uint(2), response.ID)
		assert.Equal(t, "User Name", response.Name)
	})

	t.Run("got error on GetUserByID Repository when getting user by id user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewGetUserByIdUseCase(mockUserAdminRepository)

		ctx := context.TODO()

		mockUserAdminRepository.On("GetUserById", ctx, uint(2)).
			Return(dto.UserAdmin{}, &responses.NetworkError{
				Code: 401,
			})

		response, err := sut.Execute(ctx, uint(2))

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got success when login user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewLoginUserUseCase(mockUserAdminRepository)

		ctx := context.TODO()

		mockUserAdminRepository.On("Login", ctx, "12345678910").Return("TOKEN", nil)

		response, err := sut.Execute(ctx, "12345678910")

		assert.NoError(t, err)
		assert.Equal(t, "TOKEN", response.AccessToken)
	})

	t.Run("got error on Login Repository when login user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewLoginUserUseCase(mockUserAdminRepository)

		ctx := context.TODO()

		mockUserAdminRepository.On("Login", ctx, "12345678910").Return("", &responses.NetworkError{
			Code: 404,
		})

		response, err := sut.Execute(ctx, "12345678910")

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got success when update user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewUpdateUserUseCase(NewValidateCPFUseCase(), mockUserAdminRepository)

		ctx := context.TODO()

		mockUserAdminRepository.On("UpdateUser", ctx, mockUserAdmin()).Return(nil)

		err := sut.Execute(ctx, newUserAdmin())

		assert.NoError(t, err)
	})

	t.Run("got error on UpdateUser Repository when update user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewUpdateUserUseCase(NewValidateCPFUseCase(), mockUserAdminRepository)

		ctx := context.TODO()

		mockUserAdminRepository.On("UpdateUser", ctx, mockUserAdmin()).Return(&responses.NetworkError{
			Code: 500,
		})

		err := sut.Execute(ctx, newUserAdmin())

		assert.Error(t, err)
	})

	t.Run("got error on invalid cpf when update user admin use case", func(t *testing.T) {
		t.Parallel()

		mockUserAdminRepository := new(MockUserAdminRepository)
		sut := NewUpdateUserUseCase(NewValidateCPFUseCase(), mockUserAdminRepository)

		ctx := context.TODO()

		err := sut.Execute(ctx, newInvalidUserAdmin())

		assert.Error(t, err)
	})
}

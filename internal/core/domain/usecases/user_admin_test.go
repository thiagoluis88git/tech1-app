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
}

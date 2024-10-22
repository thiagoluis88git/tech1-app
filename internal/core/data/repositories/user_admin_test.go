package repositories_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/data/model"
	"github.com/thiagoluis88git/tech1/internal/core/data/repositories"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/pkg/database"
)

const (
	insertQuery = "INSERT INTO `user_admins` (`created_at`,`updated_at`,`deleted_at`,`name`,`cpf`,`email`) VALUES (?,?,?,?,?,?)"
)

func mockDTOUserAdmin() dto.UserAdmin {
	return dto.UserAdmin{
		Name:  "NAME",
		CPF:   "CPF",
		Email: "EMAIL",
	}
}

func mockModelUserAdmin() *model.UserAdmin {
	return &model.UserAdmin{
		Name:  "NAME",
		CPF:   "CPF",
		Email: "EMAIL",
	}
}

func TestUserAdminLocal(t *testing.T) {
	t.Parallel()

	t.Run("got success when saving user admin local", func(t *testing.T) {
		t.Parallel()

		db, sqlMock, err := SetupDBMocks()

		assert.NoError(t, err)

		sqlMock.ExpectBegin()
		sqlMock.ExpectExec(insertQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "NAME", "CPF", "EMAIL").
			WillReturnResult(sqlmock.NewResult(1, 1))
		sqlMock.ExpectCommit()

		cognitoRemote := new(MockCognitoRemoteDataSource)
		localDs := repositories.NewUserAdminRepository(&database.Database{Connection: db}, cognitoRemote)

		cognitoRemote.On("SignUpAdmin", mockModelUserAdmin()).Return(nil)

		id, err := localDs.CreateUser(context.TODO(), mockDTOUserAdmin())

		assert.NoError(t, err)
		assert.Equal(t, uint(1), id)
	})
}

package repositories_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1/internal/core/data/model"
	"github.com/thiagoluis88git/tech1/internal/core/data/repositories"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/pkg/database"
	"github.com/thiagoluis88git/tech1/pkg/responses"
	"gorm.io/gorm"
)

const (
	insertQuery      = "INSERT INTO `user_admins` (`created_at`,`updated_at`,`deleted_at`,`name`,`cpf`,`email`) VALUES (?,?,?,?,?,?)"
	selectQueryByID  = "SELECT * FROM `user_admins` WHERE `user_admins`.`id` = ? AND `user_admins`.`deleted_at` IS NULL ORDER BY `user_admins`.`id` LIMIT ?"
	selectQueryByCPF = "SELECT * FROM `user_admins` WHERE cpf = ? AND `user_admins`.`deleted_at` IS NULL ORDER BY `user_admins`.`id` LIMIT ?"
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

	t.Run("got error on Cognito remote when saving user admin local", func(t *testing.T) {
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

		cognitoRemote.On("SignUpAdmin", mockModelUserAdmin()).Return(&responses.NetworkError{
			Code: 400,
		})

		id, err := localDs.CreateUser(context.TODO(), mockDTOUserAdmin())

		assert.Error(t, err)
		assert.Equal(t, uint(0), id)
	})

	t.Run("got error on Create User DB when saving user admin local", func(t *testing.T) {
		t.Parallel()

		db, sqlMock, err := SetupDBMocks()

		assert.NoError(t, err)

		sqlMock.ExpectBegin()
		sqlMock.ExpectExec(insertQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "NAME", "CPF", "EMAIL").
			WillReturnError(errors.New("Error on DB"))
		sqlMock.ExpectCommit()

		cognitoRemote := new(MockCognitoRemoteDataSource)
		localDs := repositories.NewUserAdminRepository(&database.Database{Connection: db}, cognitoRemote)

		cognitoRemote.On("SignUpAdmin", mockModelUserAdmin()).Return(nil)

		id, err := localDs.CreateUser(context.TODO(), mockDTOUserAdmin())

		assert.Error(t, err)
		assert.Equal(t, uint(0), id)
	})

	t.Run("got success when updating user admin local", func(t *testing.T) {
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

		err = localDs.UpdateUser(context.TODO(), mockDTOUserAdmin())

		assert.NoError(t, err)
	})

	t.Run("got error on Update User DB when updating user admin local", func(t *testing.T) {
		t.Parallel()

		db, sqlMock, err := SetupDBMocks()

		assert.NoError(t, err)

		sqlMock.ExpectBegin()
		sqlMock.ExpectExec(insertQuery).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "NAME", "CPF", "EMAIL").
			WillReturnError(errors.New("Error on DB"))
		sqlMock.ExpectCommit()

		cognitoRemote := new(MockCognitoRemoteDataSource)
		localDs := repositories.NewUserAdminRepository(&database.Database{Connection: db}, cognitoRemote)

		err = localDs.UpdateUser(context.TODO(), mockDTOUserAdmin())

		assert.Error(t, err)
	})

	t.Run("got success when getting user admin by id local", func(t *testing.T) {
		t.Parallel()

		db, sqlMock, err := SetupDBMocks()

		assert.NoError(t, err)

		sqlMock.ExpectQuery(selectQueryByID).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"name", "cpf", "email"}).AddRow("Name", "CPF", "Email"))

		cognitoRemote := new(MockCognitoRemoteDataSource)
		localDs := repositories.NewUserAdminRepository(&database.Database{Connection: db}, cognitoRemote)

		userAdmin, err := localDs.GetUserById(context.TODO(), uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, userAdmin)
	})

	t.Run("got error when getting user admin by id local", func(t *testing.T) {
		t.Parallel()

		db, sqlMock, err := SetupDBMocks()

		assert.NoError(t, err)

		sqlMock.ExpectQuery(selectQueryByID).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(gorm.ErrRecordNotFound)

		cognitoRemote := new(MockCognitoRemoteDataSource)
		localDs := repositories.NewUserAdminRepository(&database.Database{Connection: db}, cognitoRemote)

		userAdmin, err := localDs.GetUserById(context.TODO(), uint(1))

		assert.Error(t, err)
		assert.Empty(t, userAdmin)
	})

	t.Run("got success when getting user admin by cpf local", func(t *testing.T) {
		t.Parallel()

		db, sqlMock, err := SetupDBMocks()

		assert.NoError(t, err)

		sqlMock.ExpectQuery(selectQueryByCPF).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"name", "cpf", "email"}).AddRow("Name", "CPF", "Email"))

		cognitoRemote := new(MockCognitoRemoteDataSource)
		localDs := repositories.NewUserAdminRepository(&database.Database{Connection: db}, cognitoRemote)

		userAdmin, err := localDs.GetUserByCPF(context.TODO(), "CPF")

		assert.NoError(t, err)
		assert.NotEmpty(t, userAdmin)
	})

	t.Run("got error when getting user admin by cpf local", func(t *testing.T) {
		t.Parallel()

		db, sqlMock, err := SetupDBMocks()

		assert.NoError(t, err)

		sqlMock.ExpectQuery(selectQueryByCPF).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(gorm.ErrRecordNotFound)

		cognitoRemote := new(MockCognitoRemoteDataSource)
		localDs := repositories.NewUserAdminRepository(&database.Database{Connection: db}, cognitoRemote)

		userAdmin, err := localDs.GetUserByCPF(context.TODO(), "CPF")

		assert.Error(t, err)
		assert.Empty(t, userAdmin)
	})

	t.Run("got success when login user admin local", func(t *testing.T) {
		t.Parallel()

		db, _, err := SetupDBMocks()

		assert.NoError(t, err)

		cognitoRemote := new(MockCognitoRemoteDataSource)
		localDs := repositories.NewUserAdminRepository(&database.Database{Connection: db}, cognitoRemote)

		cognitoRemote.On("Login", "12345678910").Return("TOKEN", nil)

		token, err := localDs.Login(context.TODO(), "12345678910")

		assert.NoError(t, err)
		assert.Equal(t, "TOKEN", token)
	})

	t.Run("got error when login user admin local", func(t *testing.T) {
		t.Parallel()

		db, _, err := SetupDBMocks()

		assert.NoError(t, err)

		cognitoRemote := new(MockCognitoRemoteDataSource)
		localDs := repositories.NewUserAdminRepository(&database.Database{Connection: db}, cognitoRemote)

		cognitoRemote.On("Login", "12345678910").Return("", &responses.NetworkError{
			Code: 400,
		})

		token, err := localDs.Login(context.TODO(), "12345678910")

		assert.Error(t, err)
		assert.Empty(t, token)
	})
}

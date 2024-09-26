package repositories

import (
	"context"

	"github.com/thiagoluis88git/tech1/internal/core/data/model"
	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/domain/repository"
	"github.com/thiagoluis88git/tech1/internal/integrations/remote"
	"github.com/thiagoluis88git/tech1/pkg/responses"

	"gorm.io/gorm"
)

type UserAdminRepository struct {
	db            *gorm.DB
	cognitoRemote remote.CognitoRemoteDataSource
}

func NewUserAdminRepository(db *gorm.DB, cognitoRemote remote.CognitoRemoteDataSource) repository.UserAdminRepository {
	return &UserAdminRepository{
		db:            db,
		cognitoRemote: cognitoRemote,
	}
}

func (repository *UserAdminRepository) CreateUser(ctx context.Context, customer dto.UserAdmin) (uint, error) {
	userEntity := &model.UserAdmin{
		Name:  customer.Name,
		CPF:   customer.CPF,
		Email: customer.Email,
	}

	err := repository.cognitoRemote.SignUpAdmin(userEntity)

	if err != nil {
		return 0, responses.GetCognitoError(err)
	}

	err = repository.db.WithContext(ctx).Create(userEntity).Error

	if err != nil {
		return 0, responses.GetDatabaseError(err)
	}

	return userEntity.ID, nil
}

func (repository *UserAdminRepository) UpdateUser(ctx context.Context, customer dto.UserAdmin) error {
	userEntity := &model.UserAdmin{
		Model: gorm.Model{ID: customer.ID},
		Name:  customer.Name,
		CPF:   customer.CPF,
		Email: customer.Email,
	}

	err := repository.db.WithContext(ctx).Save(&userEntity).Error

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}

func (repository *UserAdminRepository) GetUserById(ctx context.Context, id uint) (dto.UserAdmin, error) {
	var userEntity model.UserAdmin

	err := repository.
		db.WithContext(ctx).
		First(&userEntity, id).
		Error

	if err != nil {
		return dto.UserAdmin{}, responses.GetDatabaseError(err)
	}

	return repository.populateUser(userEntity), nil
}

func (repository *UserAdminRepository) GetUserByCPF(ctx context.Context, cpf string) (dto.UserAdmin, error) {
	var userEntity model.UserAdmin

	err := repository.
		db.WithContext(ctx).
		Where("cpf = ?", cpf).
		First(&userEntity).
		Error

	if err != nil {
		return dto.UserAdmin{}, responses.GetDatabaseError(err)
	}

	return repository.populateUser(userEntity), nil
}

func (repository *UserAdminRepository) populateUser(userrEntity model.UserAdmin) dto.UserAdmin {
	return dto.UserAdmin{
		ID:    userrEntity.ID,
		Name:  userrEntity.Name,
		CPF:   userrEntity.CPF,
		Email: userrEntity.Email,
	}
}

func (repository *UserAdminRepository) Login(ctx context.Context, cpf string) (string, error) {
	token, err := repository.cognitoRemote.Login(cpf)

	if err != nil {
		return "", responses.GetDatabaseError(err)
	}

	return token, nil
}

package usecases

import (
	"context"
	"net/http"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/domain/repository"
	"github.com/thiagoluis88git/tech1/pkg/responses"
)

type CreateUserUseCase struct {
	validateCPFUseCase *ValidateCPFUseCase
	repository         repository.UserAdminRepository
}

type UpdateUserUseCase struct {
	validateCPFUseCase *ValidateCPFUseCase
	repository         repository.UserAdminRepository
}

type GetUserByCPFUseCase struct {
	validateCPFUseCase *ValidateCPFUseCase
	repository         repository.UserAdminRepository
}

type GetUserByIdUseCase struct {
	repository repository.UserAdminRepository
}

type LoginUserUseCase struct {
	repository repository.UserAdminRepository
}

func NewUpdateUserUseCase(validateCPFUseCase *ValidateCPFUseCase, repository repository.UserAdminRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		validateCPFUseCase: validateCPFUseCase,
		repository:         repository,
	}
}

func NewCreateUserUseCase(validateCPFUseCase *ValidateCPFUseCase, repository repository.UserAdminRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		validateCPFUseCase: validateCPFUseCase,
		repository:         repository,
	}
}

func NewGetUserByCPFUseCase(validateCPFUseCase *ValidateCPFUseCase, repository repository.UserAdminRepository) *GetUserByCPFUseCase {
	return &GetUserByCPFUseCase{
		validateCPFUseCase: validateCPFUseCase,
		repository:         repository,
	}
}

func NewGetUserByIdUseCase(repository repository.UserAdminRepository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		repository: repository,
	}
}

func NewLoginUserUseCase(repository repository.UserAdminRepository) *LoginUserUseCase {
	return &LoginUserUseCase{
		repository: repository,
	}
}

func (service *CreateUserUseCase) Execute(ctx context.Context, user dto.UserAdmin) (dto.UserAdminResponse, error) {
	cleanedCPF, validate := service.validateCPFUseCase.Execute(user.CPF)

	if !validate {
		return dto.UserAdminResponse{}, &responses.BusinessResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid CPF",
		}
	}

	user.CPF = cleanedCPF
	customerId, err := service.repository.CreateUser(ctx, user)

	if err != nil {
		return dto.UserAdminResponse{}, responses.GetResponseError(err, "UserService")
	}

	return dto.UserAdminResponse{
		Id: customerId,
	}, nil
}

func (service *UpdateUserUseCase) Execute(ctx context.Context, user dto.UserAdmin) error {
	cleanedCPF, validate := service.validateCPFUseCase.Execute(user.CPF)

	if !validate {
		return &responses.BusinessResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid CPF",
		}
	}

	user.CPF = cleanedCPF
	err := service.repository.UpdateUser(ctx, user)

	if err != nil {
		return responses.GetResponseError(err, "UserService")
	}

	return nil
}

func (service *GetUserByIdUseCase) Execute(ctx context.Context, id uint) (dto.UserAdmin, error) {
	user, err := service.repository.GetUserById(ctx, id)

	if err != nil {
		return dto.UserAdmin{}, responses.GetResponseError(err, "UserService")
	}

	return user, nil
}

func (service *GetUserByCPFUseCase) Execute(ctx context.Context, cpf string) (dto.UserAdmin, error) {
	cleanedCPF, validate := service.validateCPFUseCase.Execute(cpf)

	if !validate {
		return dto.UserAdmin{}, &responses.BusinessResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid CPF",
		}
	}

	user, err := service.repository.GetUserByCPF(ctx, cleanedCPF)

	if err != nil {
		return dto.UserAdmin{}, responses.GetResponseError(err, "UserService")
	}

	return user, nil
}

func (uc *LoginUserUseCase) Execute(ctx context.Context, cpf string) (dto.Token, error) {
	token, err := uc.repository.Login(ctx, cpf)

	if err != nil {
		return dto.Token{}, responses.GetResponseError(err, "UserService")
	}

	return dto.Token{
		AccessToken: token,
	}, nil
}

package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/thiagoluis88git/tech1/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1/internal/core/domain/usecases"
	"github.com/thiagoluis88git/tech1/pkg/httpserver"
)

// @Summary Create new user admin
// @Description Create new customer. This process is not required to make an order
// @Tags UserAdmin
// @Accept json
// @Produce json
// @Param product body dto.UserAdmin true "user admin"
// @Success 200 {object} dto.UserAdminResponse
// @Failure 400 "Customer has required fields"
// @Failure 409 "This user is already added"
// @Router /auth/admin/signup [post]
func CreateUserHandler(createUserAdmin *usecases.CreateUserUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user dto.UserAdmin

		err := httpserver.DecodeJSONBody(w, r, &user)

		if err != nil {
			log.Print("decoding user body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		response, err := createUserAdmin.Execute(r.Context(), user)

		if err != nil {
			log.Print("create user", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, response)
	}
}

// @Summary Update user
// @Description Update user
// @Tags UserAdmin
// @Accept json
// @Produce json
// @Param id path int true "12"
// @Param product body dto.Customer true "customer"
// @Success 204
// @Failure 400 "User has required fields"
// @Failure 404 "User not found"
// @Router /api/users/{id} [put]
func UpdateUserHandler(updateUser *usecases.UpdateUserUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("update user", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		userId, err := strconv.Atoi(userIdStr)

		if err != nil {
			log.Print("update user", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		var user dto.UserAdmin
		err = httpserver.DecodeJSONBody(w, r, &user)

		if err != nil {
			log.Print("decoding user body for update", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		user.ID = uint(userId)
		err = updateUser.Execute(r.Context(), user)

		if err != nil {
			log.Print("update user", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseNoContentSuccess(w)
	}
}

// @Summary Get user by ID
// @Description Get user by ID
// @Tags UserAdmin
// @Accept json
// @Produce json
// @Param Id path string true "12"
// @Success 200 {object} dto.UserAdmin
// @Failure 404 "User not found"
// @Router /api/users/{id} [get]
func GetUserByIdHandler(getUserById *usecases.GetUserByIdUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdStr, err := httpserver.GetPathParamFromRequest(r, "id")

		if err != nil {
			log.Print("get user by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		userId, err := strconv.Atoi(userIdStr)

		if err != nil {
			log.Print("get user by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		user, err := getUserById.Execute(r.Context(), uint(userId))

		if err != nil {
			log.Print("get user by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, user)
	}
}

// @Summary Get user by CPF
// @Description Get user by CPF. This Endpoint can be used as a Login
// @Tags UserAdmin
// @Accept json
// @Produce json
// @Param user body dto.UserAdminForm true "UserAdminForm"
// @Success 200 {object} dto.UserAdmin
// @Failure 404 "User not found"
// @Router /api/users/login [post]
func GetUserByCPFHandler(getUserByCPF *usecases.GetUserByCPFUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userForm dto.UserAdminForm

		err := httpserver.DecodeJSONBody(w, r, &userForm)

		if err != nil {
			log.Print("decoding user form body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		user, err := getUserByCPF.Execute(r.Context(), userForm.CPF)

		if err != nil {
			log.Print("get user by id", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, user)
	}
}

// @Summary Login
// @Description Login the user by its CPF
// @Tags UserAdmin
// @Accept json
// @Produce json
// @Param customer body dto.UserAdminForm true "user form"
// @Success 200 {object} dto.Token
// @Failure 404 "User not found"
// @Router /auth/admin/login [post]
func LoginUserHandler(loginUserUseCase *usecases.LoginUserUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userForm dto.UserAdminForm

		err := httpserver.DecodeJSONBody(w, r, &userForm)

		if err != nil {
			log.Print("decoding user form body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		token, err := loginUserUseCase.Execute(r.Context(), userForm.CPF)

		if err != nil {
			log.Print("login user", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseSuccess(w, token)
	}
}

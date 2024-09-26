package dto

type UserAdmin struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" validate:"required"`
	CPF   string `json:"cpf" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UserAdminForm struct {
	CPF string `json:"cpf" validate:"required"`
}

type UserAdminResponse struct {
	Id uint `json:"id"`
}

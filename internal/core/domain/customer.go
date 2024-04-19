package domain

type Customer struct {
	Name  string `json:"name" validate:"required"`
	CPF   string `json:"cpf" validate:"required"`
	Email string `json:"email" validate:"required"`
}

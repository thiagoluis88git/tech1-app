package entities

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name  string
	CPF   string
	Email string
}

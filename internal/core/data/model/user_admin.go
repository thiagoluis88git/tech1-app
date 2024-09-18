package model

import "gorm.io/gorm"

type UserAdmin struct {
	gorm.Model
	Name  string
	CPF   string `gorm:"index;unique"`
	Email string `gorm:"unique"`
}

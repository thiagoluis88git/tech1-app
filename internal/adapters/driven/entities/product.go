package entities

import "gorm.io/gorm"

const (
	CategorySnack    = "Lanche"
	CategoryBeverage = "Bebida"
	CategoryDesert   = "Sobremesa"
	CategoryToppings = "Acompanhamento"
	CategoryCombo    = "Combo"
)

type Product struct {
	gorm.Model
	Name         string `gorm:"unique"`
	Description  string
	Category     string
	Price        int
	ProductImage []ProductImage
	ProductCombo *[]ProductCombo
}

type ProductImage struct {
	gorm.Model
	ProductID uint
	ImageUrl  string
}

type ProductCombo struct {
	gorm.Model
	ProductID   uint
	Name        string
	Description string
	Category    string
}

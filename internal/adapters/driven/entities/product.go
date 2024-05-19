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
	Name          string `gorm:"unique"`
	Description   string
	Category      string
	Price         int
	ProductImage  []ProductImage
	ComboProducts *[]ComboProduct
}

type ProductImage struct {
	gorm.Model
	ProductID uint
	ImageUrl  string
}

type ComboProduct struct {
	gorm.Model
	ProductID uint
	ComboID   uint
}

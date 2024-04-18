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
	Name         string
	Description  string
	Category     string
	Price        int
	ProducImage  []ProducImage
	ProductCombo *[]ProductCombo
}

type ProducImage struct {
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

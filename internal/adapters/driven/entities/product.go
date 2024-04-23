package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string `gorm:"unique"`
	Description  string
	Category     string
	Price        int
	ProductImage []ProductImage
}

type ProductImage struct {
	gorm.Model
	ProductID uint
	ImageUrl  string
}

type Combo struct {
	gorm.Model
	Name        string
	Description string
	Category    string
	Products    []ComboProduct
}

type ComboProduct struct {
	gorm.Model
	ProductID uint
	ComboID   uint
}

package domain

import "gorm.io/gorm"

const (
	CategorySnack    = "Snack"
	CategoryBeverage = "Beverage"
	CategoryDesert   = "Desert"
	CategoryToppings = "Toppings"
)

type ProductEntity struct {
	gorm.Model
	Id          int32
	Name        string
	Description string
	Category    string
	Price       int
	Images      []ProducImageEntity
}

type ProducImageEntity struct {
	Id        int64
	ProductId int32
	ImageUrl  string
}

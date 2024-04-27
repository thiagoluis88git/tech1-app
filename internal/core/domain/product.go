package domain

const (
	CategorySnack    = "Lanche"
	CategoryBeverage = "Bebida"
	CategoryDesert   = "Sobremesa"
	CategoryToppings = "Acompanhamento"
	CategoryCombo    = "Combo"
)

type Product struct {
	Id               uint          `json:"id"`
	Name             string        `json:"name" validate:"required"`
	Description      string        `json:"description" validate:"required"`
	Category         string        `json:"category" validate:"required"`
	Price            int           `json:"price" validate:"required"`
	Images           []ProducImage `json:"images" validate:"required"`
	ComboProductsIds *[]uint       `json:"comboProductsIds"`
}

type ProducImage struct {
	ImageUrl string `json:"imageUrl" validate:"required"`
}

type ComboForm struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Products    []uint `json:"products" validate:"required"`
}

type Combo struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Products    []Product `json:"products"`
}

package domain

type Product struct {
	Id          int32
	Name        string        `json:"name" validate:"required"`
	Description string        `json:"description" validate:"required"`
	Category    string        `json:"category" validate:"required"`
	Price       int           `json:"price" validate:"required"`
	Images      []ProducImage `json:"images" validate:"required"`
}

type ProducImage struct {
	ImageUrl string `json:"imageUrl" validate:"required"`
}

package domain

type Product struct {
	Id          int32
	Name        string
	Description string
	Category    string
	Price       int
	Images      []ProducImage
}

type ProducImage struct {
	Id        int64
	ProductId int32
	ImageUrl  string
}

package models

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	Category    *Category
}

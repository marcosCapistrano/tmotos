package models

type Repository interface {
	SearchProducts(url string) []Product
	Seed()
}

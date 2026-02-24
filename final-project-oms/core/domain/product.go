package domain

type ProductID uint
type Price uint

type Product struct {
	ID    ProductID
	Name  string
	Price Price
}

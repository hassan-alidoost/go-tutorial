package domains

type ProductID uint

type Product struct {
	ID    ProductID
	Name  string
	Price float64
	Stock uint8
}

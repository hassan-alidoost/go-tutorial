package main

import "fmt"

type ProductFilter struct {
	MinPrice  *float64
	MaxPrice  *float64
	Category  *string
	Available *bool
}

type Product struct {
	Name      string
	Price     float64
	Category  string
	Available bool
}

func filterProducts(products []Product, filter ProductFilter) []Product {
	var filteredProducts []Product

	for _, product := range products {
		if filter.MinPrice != nil && product.Price < *filter.MinPrice {
			continue
		}

		if filter.MaxPrice != nil && product.Price > *filter.MaxPrice {
			continue
		}

		if filter.Category != nil && product.Category != *filter.Category {
			continue
		}

		if filter.Available != nil && product.Available != *filter.Available {
			continue
		}

		filteredProducts = append(filteredProducts, product)
	}

	return filteredProducts
}

func main() {
	products := []Product{
		{"Laptop", 999.99, "Electronics", true},
		{"Mouse", 29.99, "Electronics", true},
		{"Desk", 299.99, "Furniture", false},
		{"Chair", 199.99, "Furniture", true},
	}

	electronicFilter := ProductFilter{
		Category: StringPtr("Electronics"),
	}
	results := filterProducts(products, electronicFilter)
	fmt.Printf("Electronics: %d products\n", len(results))
}

func StringPtr(s string) *string {
	return &s
}

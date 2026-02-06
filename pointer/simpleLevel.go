package main

import "fmt"

type Product struct {
	ID   uint
	Name string
}

func main() {
	a, b := 2, 3
	fmt.Printf("a: %v, b: %v\n", a, b)
	swap(&a, &b)
	fmt.Printf("a: %v, b: %v\n", a, b)

	product := Product{ID: 1, Name: "Test"}
	fmt.Printf("Product: %v\n", product)
	changeName(&product, "Changed")
	fmt.Printf("Product: %v\n", product)

	newProduct := createProduct(2, "new product")
	fmt.Printf("New product: %+v", newProduct)
}

func createProduct(id uint, name string) *Product {
	return &Product{ID: id, Name: name}
}

func changeName(product *Product, name string) {
	product.Name = name
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

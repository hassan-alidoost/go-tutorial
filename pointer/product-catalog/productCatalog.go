package main

import (
	"fmt"
)

type Product struct {
	ID    uint
	Name  string
	Price float64
	Stock uint
}

type Catalog struct {
	Products map[uint]Product
}


func main() {
	catalog := NewCatalog()

	phone := Product{ID: 1, Name: "apple", Price: 100, Stock: 8}
	catalog.AddProdcut(phone)
	fmt.Printf("Catalog: %v\n", catalog)

	products := catalog.ListProduct()
	fmt.Printf("Product list: %v\n", products)

	product, ok := catalog.GetProdcut(10)

	if !ok {
		fmt.Println("Product not found!")
		return
	}

	fmt.Printf("Found product by id: %v\n", product)
}

func NewCatalog() Catalog {
	catalog := Catalog{}
	catalog.Products = make(map[uint]Product, 0)
	return catalog
}

func (c Catalog) AddProdcut(product Product) {
	c.Products[product.ID] = product
}

func (c Catalog) GetProdcut(productID uint) (Product, bool) {
	product, ok := c.Products[productID];
	return product, ok
}

func (c Catalog) ListProduct() []Product {
	var products []Product
	for _, product := range c.Products {
		products = append(products, product)
	}
	return products
}
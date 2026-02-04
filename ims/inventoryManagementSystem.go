package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID       uint16
	Name     string
	Quantity uint8
	Price    float64
}

func main() {
	inventory := make(map[int]Product)
	product1 := Product{ID: 1, Name: "mobile", Quantity: 10, Price: 1000}
	
	AddProduct(inventory, product1)
	fmt.Printf("Inventory: %v\n", inventory)

	product2 := Product{ID: 1, Name: "mobile", Quantity: 5, Price: 1000}

	AddProduct(inventory, product2)
	fmt.Printf("Inventory: %v\n", inventory)

	product3 := Product{ID: 2, Name: "laptop", Quantity: 5, Price: 4000}

	AddProduct(inventory, product3)
	fmt.Printf("Inventory: %v\n", inventory)

	// error := RemoveProduct(inventory, 1)
	// if error != nil {
	// 	fmt.Println(error.Error())
	// }
	// fmt.Printf("Inventory: %v\n", inventory)

	// error = RemoveProduct(inventory, 4)
	// if error != nil {
	// 	fmt.Println(error.Error())
	// }
	// fmt.Printf("Inventory: %v\n", inventory)

	total := CalculateTotalPrice(inventory)
	fmt.Printf("Total price: %.2f", total)
}



func AddProduct(inventory map[int]Product, product Product) error {
	
	if item, ok := inventory[int(product.ID)]; ok {
		item.Quantity += product.Quantity
		inventory[int(product.ID)] = item
		fmt.Println("Product already exists, so quantity increased!")
		return nil
	}

	inventory[int(product.ID)] = product
	fmt.Println("Product added successfully")
	return nil
}


func RemoveProduct(inventory map[int]Product, product_id uint16) error {

	if _, ok := inventory[int(product_id)]; ok {
		delete(inventory, int(product_id))
		fmt.Println("Product removed successfully")
		return nil
	}

	return errors.New("Product not exists!")
}


func CheckStock(inventory map[int]Product, product_id uint16) (int8, bool) {
	if item, ok := inventory[int(product_id)]; ok {
		return int8(item.Quantity), ok
	}

	return 0, false
}


func CalculateTotalPrice(inventory map[int]Product) float64 {
	var total float64

	for _, v := range inventory {
		total += v.Price * float64(v.Quantity)
	}

	return total
}

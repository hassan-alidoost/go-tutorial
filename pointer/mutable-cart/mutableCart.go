package main

import (
	"errors"
	"fmt"
)

type CartItem struct {
	ProductID uint
	Name      string
	Price     float64
	Quantity  uint
}

type Cart struct {
	Items []CartItem
	Total float64
}

type UpdateItemPayload struct {
	Name     *string
	Price    *float64
	Quantity *uint
}

func main() {
	var cart Cart

	sumsung := CartItem{ProductID: 1, Name: "sumsung", Price: 100.0, Quantity: 5}
	apple := CartItem{ProductID: 2, Name: "apple", Price: 500.0, Quantity: 2}
	cart.AddItem(sumsung)
	cart.AddItem(apple)
	fmt.Printf("Cart: %v\n", cart)

	err := cart.RemoveItem(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Cart: %v\n", cart)

	newQty := uint(4)
	err = cart.UpdateItem(2, UpdateItemPayload{
		Quantity: &newQty,
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Cart: %v\n", cart)

	cart.CalTotal()
	fmt.Printf("Cart: %v\n", cart)
}

func (c *Cart) CalTotal() {
	var total float64
	for _, item := range c.Items {
		total += item.Price * float64(item.Quantity)
	}

	c.Total = total
}

func (c *Cart) UpdateItem(productID uint, payload UpdateItemPayload) error {
	var cartItem CartItem
	var cartItemIndex int
	for i, item := range c.Items {
		if item.ProductID == productID {
			cartItem = item
			cartItemIndex = i
			break
		}
	}

	if cartItem.ProductID == 0 {
		return errors.New("product not exists")
	}

	if payload.Name != nil && *payload.Name != cartItem.Name {
		cartItem.Name = *payload.Name
	}

	if payload.Price != nil && *payload.Price != cartItem.Price {
		cartItem.Price = *payload.Price
	}

	if payload.Quantity != nil && *payload.Quantity != cartItem.Quantity {
		cartItem.Quantity = *payload.Quantity
	}

	c.Items[cartItemIndex] = cartItem
	return nil
}

func (c *Cart) RemoveItem(productID uint) error {
	for i, item := range c.Items {
		if item.ProductID == productID {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			return nil
		}
	}

	return errors.New("product not exists")
}

func (c *Cart) AddItem(item CartItem) {
	c.Items = append(c.Items, item)
	c.Total = float64(item.Quantity) * item.Price
}

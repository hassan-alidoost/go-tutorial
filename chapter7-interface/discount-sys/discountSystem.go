package main

import "fmt"

type Price float64

type Discounter interface {
	ApplyDiscount(price Price) Price
}

type FixedDiscount struct {
	Amount Price
}

func (fd FixedDiscount) ApplyDiscount(price Price) Price {
	final := price - fd.Amount
	if final < 0 {
		return 0
	}

	return final
}

type PercentageDiscount struct {
	Percent float64
}

func (pd PercentageDiscount) ApplyDiscount(price Price) Price {
	return price - price*Price(pd.Percent)/100
}

type NoDiscount struct{}

func (nd NoDiscount) ApplyDiscount(price Price) Price {
	return price
}

type Product struct {
	Name      string
	BasePrice Price
}

func (p Product) CalFinalPrice(discounter Discounter) Price {
	return discounter.ApplyDiscount(p.BasePrice)
}

func processProductsPrice(products []Product, discounter Discounter) Price {
	var totalPrice Price
	for _, product := range products {
		totalPrice += product.CalFinalPrice(discounter)
	}
	return totalPrice
}

func main() {
	products := []Product{
		{Name: "Laptop", BasePrice: 999.99},
		{Name: "Mouse", BasePrice: 29.99},
		{Name: "Keyboard", BasePrice: 79.99},
	}

	percentage := PercentageDiscount{Percent: 10}
	fixed := FixedDiscount{Amount: 20}
	noDiscount := NoDiscount{}

	fmt.Println("10% Discount Total:", processProductsPrice(products, percentage))
	fmt.Println("$20 Discount Total:", processProductsPrice(products, fixed))
	fmt.Println("No Discount Total:", processProductsPrice(products, noDiscount))
}

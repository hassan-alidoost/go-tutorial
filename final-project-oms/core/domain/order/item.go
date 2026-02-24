package order

import (
	"errors"

	domain "github.com/hassan-alidoost/go-tutorial/final-project-oms/core/domain"
	base "github.com/hassan-alidoost/go-tutorial/final-project-oms/core/domain/base"
)

type Item struct {
	Product  domain.Product
	Quantity uint
}

func NewOrderItem(product domain.Product, qty uint) (*Item, error) {
	if qty <= 0 {
		return nil, errors.New("qty can not be less than zero")
	}
	return &Item{Product: product, Quantity: qty}, nil
}

func (oi Item) TotalPrice() base.Price {
	return base.Price(oi.Quantity) * oi.Product.Price
}

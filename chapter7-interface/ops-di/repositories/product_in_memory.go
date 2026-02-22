package repositories

import (
	"errors"
	"github.com/hassan-alidoost/go-tutorial/chapter7-interface/ops-di/domains"
)

type InMemoryProductRepo struct {
	products map[domains.ProductID]*domains.Product
}

func NewInMemoryProductRepo() *InMemoryProductRepo {
	return &InMemoryProductRepo{products: make(map[domains.ProductID]*domains.Product)}
}

func (pr *InMemoryProductRepo) Find(productID domains.ProductID) (*domains.Product, error) {
	if _, ok := pr.products[productID]; ok {
		return pr.products[productID], nil
	}

	return nil, errors.New("product does not exist")
}

func (pr *InMemoryProductRepo) UpdateStock(productID domains.ProductID, newStock uint) error {
	product, ok := pr.products[productID]
	if !ok {
		return errors.New("product does not exist")
	}

	product.Stock = uint8(newStock)
	return nil
}

func (pr *InMemoryProductRepo) Save(product domains.Product) error {
	if _, ok := pr.products[product.ID]; ok {
		return errors.New("product already exists")
	}

	pr.products[product.ID] = &product
	return nil
}

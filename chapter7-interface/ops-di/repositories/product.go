package repositories

import "github.com/hassan-alidoost/go-tutorial/chapter7-interface/ops-di/domains"

type ProductRepository interface {
	Save(product domains.Product) error
	Find(productId domains.ProductID) (*domains.Product, error)
	UpdateStock(productID domains.ProductID, newStock uint) error
}

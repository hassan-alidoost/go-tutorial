package services

import "github.com/hassan-alidoost/go-tutorial/chapter7-interface/ops-di/domains"

type InventoryService interface {
	Reserve(productId domains.ProductID, qty uint) error
	Release(productId domains.ProductID, qty uint) error
}

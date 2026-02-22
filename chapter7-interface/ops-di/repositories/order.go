package repositories

import "github.com/hassan-alidoost/go-tutorial/chapter7-interface/ops-di/domains"

type OrderRepository interface {
	Save(order domains.Order) error
	Find(orderID domains.OrderID) (*domains.Order, error)
	UpdateStatus(orderID domains.OrderID, newStatus domains.OrderStatus) error
}

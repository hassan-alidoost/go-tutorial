package repositories

import (
	"errors"
	"github.com/hassan-alidoost/go-tutorial/chapter7-interface/ops-di/domains"
)

type InMemoryOrderRepo struct {
	orders map[domains.OrderID]*domains.Order
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{orders: make(map[domains.OrderID]*domains.Order)}
}

func (or *InMemoryOrderRepo) Find(orderID domains.OrderID) (*domains.Order, error) {
	if _, ok := or.orders[orderID]; ok {
		return or.orders[orderID], nil
	}

	return nil, errors.New("order does not exist")
}

func (or *InMemoryOrderRepo) UpdateStatus(orderID domains.OrderID, newStatus domains.OrderStatus) error {
	order, ok := or.orders[orderID]
	if !ok {
		return errors.New("order does not exist")
	}

	order.Status = newStatus
	return nil
}

func (or *InMemoryOrderRepo) Save(order domains.Order) error {
	if _, ok := or.orders[order.ID]; ok {
		return errors.New("order already exists")
	}

	or.orders[order.ID] = &order
	return nil
}

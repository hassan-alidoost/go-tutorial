package adaptor

import (
	"errors"
	"fmt"

	base "github.com/hassan-alidoost/go-tutorial/final-project-oms/domain/base"
	domain "github.com/hassan-alidoost/go-tutorial/final-project-oms/domain/order"
)

type InMemoryOrderRepo struct {
	orders map[base.ID]*domain.Order
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		orders: make(map[base.ID]*domain.Order),
	}
}

func (r *InMemoryOrderRepo) Save(order *domain.Order) error {
	if order == nil {
		return errors.New("cannot save a nil order")
	}

	r.orders[order.ID] = order
	return nil
}

func (r *InMemoryOrderRepo) FindByID(ID base.ID) (*domain.Order, error) {
	if order, ok := r.orders[ID]; ok {

		return order, nil
	}

	return nil, fmt.Errorf("order with ID %d not found", ID)
}

func (r *InMemoryOrderRepo) List() []domain.Order {
	list := make([]domain.Order, 0, len(r.orders))
	for _, o := range r.orders {
		list = append(list, *o)
	}
	return list
}

func (r *InMemoryOrderRepo) Delete(ID base.ID) error {
	if _, ok := r.orders[ID]; !ok {
		return fmt.Errorf("order %d not found", ID)
	}

	delete(r.orders, ID)
	return nil
}

func (r *InMemoryOrderRepo) Clear() {
	r.orders = make(map[base.ID]*domain.Order)
}

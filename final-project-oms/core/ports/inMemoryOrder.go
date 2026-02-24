package ports

import (
	base "github.com/hassan-alidoost/go-tutorial/final-project-oms/core/domain/base"
	domain "github.com/hassan-alidoost/go-tutorial/final-project-oms/core/domain/order"
)

type OrderRepository interface {
	Save(order *domain.Order) error
	FindByID(id base.ID) (*domain.Order, error)
	List() []domain.Order
	Delete(id base.ID) error
	Clear()
}

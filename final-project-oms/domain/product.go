package domain

import (
	base "github.com/hassan-alidoost/go-tutorial/final-project-oms/domain/base"
)

type Product struct {
	ID    base.ID
	Name  string
	Price base.Price
	base.Timestamps
}

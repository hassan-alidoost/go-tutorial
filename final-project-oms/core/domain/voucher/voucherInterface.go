package voucher

import domain "github.com/hassan-alidoost/go-tutorial/final-project-oms/core/domain/base"

type Voucher interface {
	Apply(total domain.Price) domain.Price
	Code() string
}

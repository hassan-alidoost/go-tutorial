package voucher

import (
	"github.com/hassan-alidoost/go-tutorial/final-project-oms/domain/base"
)

type Voucher interface {
	Apply(total domain.Price) domain.Price
	Code() string
}

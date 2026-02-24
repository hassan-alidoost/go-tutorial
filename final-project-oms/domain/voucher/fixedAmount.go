package voucher

import (
	"errors"

	base "github.com/hassan-alidoost/go-tutorial/final-project-oms/domain/base"
)

type FixedAmountVoucher struct {
	code     string
	amount   base.Price
	minTotal base.Price
}

func NewFixedAmountVoucher(code string, amount base.Price, minTotal base.Price) (*FixedAmountVoucher, error) {
	if amount < 0 {
		return nil, errors.New("invalid amount")
	}
	return &FixedAmountVoucher{code: code, amount: amount, minTotal: minTotal}, nil
}

func (v *FixedAmountVoucher) Apply(total base.Price) base.Price {
	if total < v.minTotal {
		return total
	}
	return total - v.amount
}

func (v *FixedAmountVoucher) Code() string {
	return v.code
}

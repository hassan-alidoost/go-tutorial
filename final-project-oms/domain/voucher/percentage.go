package voucher

import (
	"errors"

	base "github.com/hassan-alidoost/go-tutorial/final-project-oms/domain/base"
)

type PercentageVoucher struct {
	code    string
	percent base.Percent
}

func NewPercentageVoucher(code string, percent base.Percent) (*PercentageVoucher, error) {
	if percent < 0 || percent > 100 {
		return nil, errors.New("invalid percent")
	}
	return &PercentageVoucher{code: code, percent: percent}, nil
}

func (v *PercentageVoucher) Apply(total base.Price) base.Price {
	return total - (total*base.Price(v.percent))/100
}

func (v *PercentageVoucher) Code() string {
	return v.code
}

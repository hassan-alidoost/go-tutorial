package order

import (
	"errors"
	"fmt"
	"time"

	"github.com/hassan-alidoost/go-tutorial/final-project-oms/core/domain"
	base "github.com/hassan-alidoost/go-tutorial/final-project-oms/core/domain/base"
	"github.com/hassan-alidoost/go-tutorial/final-project-oms/core/domain/voucher"
)

const maxOrderItem = 10

type Order struct {
	ID      base.ID
	Items   []Item
	State   State
	Voucher voucher.Voucher
	base.Timestamps
}

func NewOrder(ID base.ID) *Order {
	return &Order{
		ID:    ID,
		Items: make([]Item, 0, maxOrderItem),
		State: Created,
		Timestamps: base.Timestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func (o *Order) AddItem(product domain.Product, qty uint) error {
	defer func() { o.UpdatedAt = time.Now() }()

	if o.State != Created {
		return fmt.Errorf("can not add item in %s state", o.State)
	}

	if len(o.Items) >= maxOrderItem {
		return errors.New("can not add item, items limit exceeded")
	}

	if qty < 0 {
		return errors.New("qty can not be less than zero")
	}

	o.Items = append(o.Items, Item{Product: product, Quantity: qty})

	return nil
}

func (o *Order) ApplyVoucher(v voucher.Voucher) error {
	defer func() { o.UpdatedAt = time.Now() }()

	if o.State != Created {
		return fmt.Errorf("can not apply voucher in %s state", o.State)
	}

	if o.Voucher != nil {
		return errors.New("voucher is already applied")
	}

	if v == nil {
		return errors.New("can not apply nil voucher")
	}

	o.Voucher = v

	return nil
}

func (o *Order) TotalAmount() (base.Price, error) {
	if len(o.Items) == 0 {
		return 0, errors.New("order items is empty")
	}

	var total base.Price
	for _, item := range o.Items {
		total += item.TotalPrice()
	}

	if o.Voucher != nil {
		total = o.Voucher.Apply(total)
	}

	if total < 0 {
		return 0, errors.New("total amount must be greater than zero")
	}

	return total, nil
}

func (o *Order) Pay() error {
	defer func() { o.UpdatedAt = time.Now() }()

	//todo-mock-pay

	if err := o.TransitionTo(Cancelled); err != nil {
		return err
	}

	return nil
}

func (o *Order) Cancel() error {
	defer func() { o.UpdatedAt = time.Now() }()

	if err := o.TransitionTo(Cancelled); err != nil {
		return err
	}

	return nil
}

func (o *Order) SnapshotItems() []Item {
	snapshot := make([]Item, len(o.Items))
	copy(snapshot, o.Items)
	return snapshot
}

package order

import (
	"errors"
	"fmt"
)

type State uint8

const (
	Created State = iota
	Paid
	VendorAccepted
	Shipped
	Delivered
	Cancelled
)

func (o *Order) TransitionTo(newState State) error {
	if o == nil {
		return errors.New("can not change state on nil")
	}

	var isValid bool
	switch o.State {
	case Cancelled:
		return errors.New("can not change state. order is cancelled")
	case Created:
		isValid = newState == Paid || newState == Cancelled
	case Paid:
		isValid = newState == VendorAccepted || newState == Cancelled
	case VendorAccepted:
		isValid = newState == Shipped || newState == Cancelled
	case Shipped:
		isValid = newState == Delivered || newState == Cancelled
	case Delivered:
		isValid = false
	}

	if !isValid {
		return fmt.Errorf("can not transition from %s to %s", o.State, newState)
	}

	o.State = newState
	return nil
}

func (os State) String() string {
	return [...]string{
		"Created",
		"Paid",
		"VendorAccepted",
		"Shipped",
		"Delivered",
		"Cancelled",
	}[os]
}

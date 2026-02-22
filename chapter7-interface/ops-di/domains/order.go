package domains

import "time"

type OrderStatus string
type OrderID uint

const (
	OrderPending   OrderStatus = "pending"
	OrderConfirmed OrderStatus = "confirmed"
	OrderShipped   OrderStatus = "shipped"
	OrderDelivered OrderStatus = "delivered"
	OrderCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID         OrderID
	CustomerID uint
	Items      []OrderItem
	Status     OrderStatus
	Total      float64
	CreatedAt  time.Time
}

type OrderItem struct {
	ProductID ProductID
	Quantity  uint
	Price     float64
}

func (o *Order) CalTotal() {
	var total float64
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	o.Total = total
}

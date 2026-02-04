package main

import "fmt"

type OrderItem struct {
	ProductID uint16
	Quantity  uint8
	Price     float64
}

type Order struct {
	ID       uint16
	Items    []OrderItem
	Discount float64
	Total    float64
	Status   string
}

type DiscountRule struct {
	MinAmount       float64
	DiscountPercent float64
	Description     string
}

func main() {
	rules := []DiscountRule{
		{MinAmount: 100.0, DiscountPercent: 5.0, Description: "5% off for orders over $100"},
		{MinAmount: 500.0, DiscountPercent: 10.0, Description: "10% off for orders over $500"},
		{MinAmount: 1000.0, DiscountPercent: 15.0, Description: "15% off for orders over $1000"},
	}

	orders := []Order{
		{
			ID: 1,
			Items: []OrderItem{
				{ProductID: 1, Quantity: 2, Price: 50.0},
				{ProductID: 2, Quantity: 1, Price: 30.0},
			},
			Status: "pending",
		},
		{
			ID: 2,
			Items: []OrderItem{
				{ProductID: 3, Quantity: 10, Price: 100.0},
			},
			Status: "pending",
		},
	}

	processedOrders := ProcessOrders(orders, rules)

	completedOrders := FilterOrdersByStatus(processedOrders, "completed")

	fmt.Printf("Complete orders: %v\n", completedOrders)

	stats := CalculateOrderStatistics(processedOrders)
	fmt.Printf("Orders statistics: %v", stats)

}

func CalculateSubtotal(order Order) float64 {
	var subTotal float64
	for _, item := range order.Items {
		subTotal += item.Price * float64(item.Quantity)
	}
	return subTotal
}

// Apply discount based on rules
func ApplyDiscountRules(order Order, rules []DiscountRule) Order {
	order.Total = CalculateSubtotal(order)
	fmt.Printf("Order subTotal: %.2f\n", order.Total)

	var selectedRule DiscountRule

	for _, rule := range rules {
		if order.Total >= rule.MinAmount && rule.MinAmount > selectedRule.MinAmount {
			selectedRule = rule
		}
	}

	if selectedRule.DiscountPercent == 0 {
		fmt.Println("No discount available for this order amount!")
	} else {
		fmt.Printf("Discount description: %s\n", selectedRule.Description)
	}

	order.Discount = order.Total * selectedRule.DiscountPercent/100
	fmt.Printf("Order discount: %.2f\n", order.Discount)

	return order
}

// Process multiple orders simultaneously
func ProcessOrders(orders []Order, rules []DiscountRule) []Order {
	for i, order := range orders {
		order.Status = "processing"
		fmt.Println("Processing order: ", order.ID)

		order = ApplyDiscountRules(order, rules)
		order.Total -= order.Discount
		fmt.Printf("Order total: %.2f\n", order.Total)

		if order.Total < 0 {
			order.Status = "cancelled"
			fmt.Println("Order with ID ", order.ID, " canceled!")
			continue
		}

		order.Status = "completed"
		fmt.Println("Order with ID ", order.ID, " completed!")

		orders[i] = order
	}
	
	return orders
}

// Filter orders by status
func FilterOrdersByStatus(orders []Order, status string) []Order {
	filteredOrders := make([]Order, 0)
	for _, order := range orders {
		if order.Status == status {
			filteredOrders = append(filteredOrders, order)
		}
	}

	return filteredOrders
}

// Calculate order statistics
func CalculateOrderStatistics(orders []Order) map[string]interface{} {
	return map[string]interface{}{
		"total_orders" : len(orders),
		"total_revenue": TotalRevenue(orders) ,
		"average_order_value": AvgOrderValue(orders),
		"completed_orders": len(FilterOrdersByStatus(orders, "completed")),
		"total_discount": TotalDiscount(orders), 
	}
}

func AvgOrderValue(orders []Order) float64 {
	var average float64

	for _, order := range orders {
		average += order.Total
	}

	return average/float64(len(orders))
}

func TotalRevenue(orders []Order) float64 {
	var totalRevenue float64

	for _, order := range orders {
		totalRevenue += order.Total
	}

	return totalRevenue
}

func TotalDiscount(orders []Order) float64  {
	var totalDiscount float64
	
	for _, order := range orders {
		totalDiscount += order.Discount
	}

	return totalDiscount
}

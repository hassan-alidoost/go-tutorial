package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	"github.com/hassan-alidoost/go-tutorial/final-project-oms/adaptor"
	"github.com/hassan-alidoost/go-tutorial/final-project-oms/domain"
	base "github.com/hassan-alidoost/go-tutorial/final-project-oms/domain/base"
	"github.com/hassan-alidoost/go-tutorial/final-project-oms/domain/order"
	voucherDomain "github.com/hassan-alidoost/go-tutorial/final-project-oms/domain/voucher"
)

func main() {
	repo := adaptor.NewInMemoryOrderRepo()

	laptop := domain.Product{
		ID:         base.ID(rand.Uint()),
		Name:       "Gaming Laptop",
		Price:      1500,
		Timestamps: base.Timestamps{CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	mouse := domain.Product{
		ID:         base.ID(rand.Uint()),
		Name:       "Wireless Mouse",
		Price:      50,
		Timestamps: base.Timestamps{CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	var orderID = base.ID(rand.Uint())
	myOrder := order.NewOrder(orderID)

	fmt.Println("--- Adding Items ---")
	if err := myOrder.AddItem(laptop, 1); err != nil {
		log.Fatalf("Failed to add %s: %v", laptop.Name, err)
	}
	if err := myOrder.AddItem(mouse, 2); err != nil {
		log.Fatalf("Failed to add %s: %v", mouse.Name, err)
	}

	items := myOrder.SnapshotItems()
	fmt.Printf("Added %d items to Order #%d\n", len(items), myOrder.ID)

	fmt.Println("\n--- Applying Voucher ---")
	voucher, err := voucherDomain.NewFixedAmountVoucher("test", 100, 500)
	if err != nil {
		fmt.Printf("Voucher error: %v\n", err)
	}

	if err = myOrder.ApplyVoucher(voucher); err != nil {
		fmt.Printf("Voucher error: %v\n", err)
	}

	total, _ := myOrder.TotalAmount()
	fmt.Printf("Total Amount after voucher: $%d\n", total)

	fmt.Println("\n--- Payment & State ---")
	if err = myOrder.Pay(); err != nil {
		fmt.Printf("Payment failed: %v\n", err)
	}
	fmt.Printf("Order State: %v\n", myOrder.State)

	fmt.Println("\n--- Repository Operations ---")
	if err = repo.Save(myOrder); err != nil {
		log.Fatalf("Save failed: %v", err)
	}

	foundOrder, err := repo.FindByID(orderID)
	if err != nil {
		log.Fatalf("Find failed: %v", err)
	}
	fmt.Printf("Successfully retrieved Order #%d from Repo\n", foundOrder.ID)

	fmt.Println("\n--- Enforcement Check ---")
	err = foundOrder.AddItem(mouse, 1)
	if err != nil {
		fmt.Printf("Expected Error caught: %v\n", err)
	}

	allOrders := repo.List()
	fmt.Printf("\nTotal orders in system: %d\n", len(allOrders))
}

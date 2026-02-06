package main

import (
	"fmt"
	"sync"
	"time"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]Order, 0, 100)
	},
}

type Order struct {
	ID    int
	Total float64
}

func getBuffer() []Order {
	return bufferPool.Get().([]Order)
}

func putBuffer(buf []Order) {
	buf = buf[:0]
	bufferPool.Put(buf)
}

func processBatch(orders []Order) float64 {
	var total float64
	for _, order := range orders {
		total += order.Total
	}
	return total
}

func processOrderEfficiently(allOrders []Order, batchSize int) {
	buffer := getBuffer()
	defer putBuffer(buffer)

	batchNum := 1

	for _, order := range allOrders {
		buffer = append(buffer, order)

		if len(buffer) >= batchSize {
			total := processBatch(buffer)
			fmt.Printf("Batch %d: total %.2f\n", batchNum, total)
			buffer = buffer[:0]
			batchNum++
		}
	}

	if len(buffer) > 0 {
		total := processBatch(buffer)
		fmt.Printf("Batch %d: total %.2f\n", batchNum, total)
	}
}

func main() {
	// Generate 1000 orders
	orders := make([]Order, 1240)
	for i := range orders {
		orders[i] = Order{
			ID:    i + 1,
			Total: float64(i%100) * 10.5,
		}
	}

	start := time.Now()
	processOrderEfficiently(orders, 100)
	fmt.Printf("Took: %v\n", time.Since(start))
}

package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

type Order struct {
	ID     uint
	Amount float64
}

type Statistics struct {
	TotalRevenue float64
	AverageOrder float64
	MaxOrder     float64
	MinOrder     float64
	OrderCount   int
}

const batchSize int = 1000

var orderPool = sync.Pool{
	New: func() any {
		return make([]Order, 0, batchSize)
	},
}

func getBuffer() []Order {
	return orderPool.Get().([]Order)
}

func putBuffer(buf []Order) {
	buf = buf[:0]
	orderPool.Put(buf)
}

func processBatch(orders []Order) Statistics {
	if len(orders) == 0 {
		return Statistics{}
	}

	stats := Statistics{
		MinOrder: math.Inf(1),
		MaxOrder: 0,
	}

	for _, order := range orders {
		stats.TotalRevenue += order.Amount
		stats.OrderCount++

		if order.Amount > stats.MaxOrder {
			stats.MaxOrder = order.Amount
		}

		if order.Amount < stats.MinOrder {
			stats.MinOrder = order.Amount
		}
	}

	stats.AverageOrder = stats.TotalRevenue / float64(stats.OrderCount)

	return stats
}

func processOrdersEfficiently(orders []Order) Statistics {
	orderBuffer := getBuffer()
	defer putBuffer(orderBuffer)
	var stats Statistics
	var remainOrderStats Statistics
	batchNum := 1
	for _, order := range orders {
		orderBuffer = append(orderBuffer, order)

		if len(orderBuffer) >= batchSize {
			stats = processBatch(orderBuffer)
			orderBuffer = orderBuffer[:0]
			batchNum++
		}
	}

	if len(orderBuffer) > 0 {
		remainOrderStats = processBatch(orderBuffer)
	}

	return Statistics{
		TotalRevenue: stats.TotalRevenue + remainOrderStats.TotalRevenue,
		OrderCount:   stats.OrderCount + remainOrderStats.OrderCount,
		MaxOrder:     math.Max(stats.MaxOrder, remainOrderStats.MaxOrder),
		MinOrder:     math.Max(stats.MinOrder, remainOrderStats.MinOrder),
		AverageOrder: (stats.TotalRevenue + remainOrderStats.TotalRevenue) / float64(stats.OrderCount + remainOrderStats.OrderCount),
	}
}

func processOrdersNaive(orders []Order) Statistics {
	var orderBuffer []Order //orderBuffer := []Order{}
	var stats Statistics
	var remainOrderStats Statistics

	for _, order := range orders {
		orderBuffer = append(orderBuffer, order)

		if len(orderBuffer) >= batchSize {
			stats = processBatch(orderBuffer)
			orderBuffer = []Order{}
		}
	}

	if len(orderBuffer) > 0 {
		remainOrderStats = processBatch(orderBuffer)
	}

	return Statistics{
		TotalRevenue: stats.TotalRevenue + remainOrderStats.TotalRevenue,
		OrderCount:   stats.OrderCount + remainOrderStats.OrderCount,
		MaxOrder:     math.Max(stats.MaxOrder, remainOrderStats.MaxOrder),
		MinOrder:     math.Max(stats.MinOrder, remainOrderStats.MinOrder),
		AverageOrder: (stats.TotalRevenue + remainOrderStats.TotalRevenue) / float64(stats.OrderCount + remainOrderStats.OrderCount),
	}
}

func printMemStats(label string) {
	var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("\n%s:\n", label)
    fmt.Printf("  Alloc: %.2f MB\n", float64(m.Alloc)/1024/1024)
    fmt.Printf("  TotalAlloc: %.2f MB\n", float64(m.TotalAlloc)/1024/1024)
    fmt.Printf("  Sys: %d MB\n", m.Sys/1024/1024)
    fmt.Printf("  NumGC: %d\n", m.NumGC)
}

func main() {
// Generate 10,000 orders
    const orderCount = 10000
    orders := make([]Order, orderCount)
    for i := range orders {
        orders[i] = Order{
            ID:     uint(i + 1),
            Amount: float64(50 + (i%100)*10),
        }
    }
    
    fmt.Printf("Processing %d orders...\n", orderCount)
    
    // Warm up GC
    runtime.GC()
    
    // Test efficient implementation
    printMemStats("Before Efficient")
    startEfficient := time.Now()
    statsEfficient := processOrdersEfficiently(orders)
    durationEfficient := time.Since(startEfficient)
    printMemStats("After Efficient")
    
    fmt.Printf("\nEfficient Results:\n")
    fmt.Printf("  Duration: %v\n", durationEfficient)
    fmt.Printf("  Total Revenue: $%.2f\n", statsEfficient.TotalRevenue)
    fmt.Printf("  Average Order: $%.2f\n", statsEfficient.AverageOrder)
    fmt.Printf("  Max Order: $%.2f\n", statsEfficient.MaxOrder)
    fmt.Printf("  Min Order: $%.2f\n", statsEfficient.MinOrder)
    
    // Force GC before naive test
    runtime.GC()
    time.Sleep(100 * time.Millisecond)
    
    // Test naive implementation
    printMemStats("Before Naive")
    startNaive := time.Now()
    statsNaive := processOrdersNaive(orders)
    durationNaive := time.Since(startNaive)
    printMemStats("After Naive")
	
    fmt.Printf("\nNaive Results:\n")
    fmt.Printf("  Duration: %v\n", durationNaive)
    fmt.Printf("  Total Revenue: $%.2f\n", statsNaive.TotalRevenue)
    fmt.Printf("  Average Order: $%.2f\n", statsNaive.AverageOrder)
    fmt.Printf("  Max Order: $%.2f\n", statsNaive.MaxOrder)
    fmt.Printf("  Min Order: $%.2f\n", statsNaive.MinOrder)
    
    // Compare
    fmt.Printf("\n=== Performance Comparison ===\n")
    speedup := float64(durationNaive) / float64(durationEfficient)
    fmt.Printf("Speedup: %.2fx faster\n", speedup)
    
    if durationEfficient < durationNaive {
        fmt.Println("✅ Efficient implementation is faster!")
    } else {
        fmt.Println("⚠️ Naive implementation was faster (unexpected)")
    }
}

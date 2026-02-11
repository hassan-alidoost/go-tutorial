package main

import "fmt"

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

type LRUCache struct {
	Capacity int
	Cache    map[int]*Node
	Head     *Node
	Tail     *Node
}

func Constructor(cap int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Prev = tail
	tail.Next = head

	return LRUCache{
		Capacity: cap,
		Cache:    make(map[int]*Node, cap),
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Cache[key]
	if !ok {
		return -1
	}

	this.removeNode(node)
	this.insertToFront(node)

	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		node.Value = value
		this.removeNode(node)
		this.insertToFront(node)
		return
	}

	if len(this.Cache) == this.Capacity {
		lru := this.Tail.Next
		this.removeNode(lru)
		delete(this.Cache, lru.Key)
	}

	node := &Node{Key: key, Value: value}
	this.Cache[key] = node
	this.insertToFront(node)
}

//Tail <---> A <---> Node <---> Head

func (this *LRUCache) insertToFront(node *Node) {
	node.Next = this.Head
	node.Prev = this.Head.Prev
	this.Head.Prev.Next = node
	this.Head.Prev = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

func (this *LRUCache) Print() {
	fmt.Print("Cache (LRU -> MRU): ")
	current := this.Tail.Next
	for current != this.Head {
		fmt.Printf("[%d=%d] ", current.Key, current.Value)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	cache := Constructor(2)

	fmt.Println("=== LRU Cache Test ===")

	cache.Put(1, 1)
	fmt.Println("Put(1, 1)")
	cache.Print()

	cache.Put(2, 2)
	fmt.Println("Put(2, 2)")
	cache.Print()

	result := cache.Get(1)
	fmt.Printf("Get(1) = %d\n", result)
	cache.Print()

	cache.Put(3, 3)
	fmt.Println("Put(3, 3) - evicts key 2")
	cache.Print()

	result = cache.Get(2)
	fmt.Printf("Get(2) = %d (not found)\n", result)

	cache.Put(4, 4)
	fmt.Println("Put(4, 4) - evicts key 1")
	cache.Print()

	result = cache.Get(1)
	fmt.Printf("Get(1) = %d (not found)\n", result)

	result = cache.Get(3)
	fmt.Printf("Get(3) = %d\n", result)
	cache.Print()

	result = cache.Get(4)
	fmt.Printf("Get(4) = %d\n", result)
	cache.Print()

	// E-Commerce Use Case
	fmt.Println("\n=== E-Commerce Product Cache ===")
	productCache := Constructor(3)

	// Cache product details
	productCache.Put(1001, 99999) // Laptop: $999.99
	productCache.Put(1002, 2999)  // Mouse: $29.99
	productCache.Put(1003, 7999)  // Keyboard: $79.99

	fmt.Println("Initial cache (3 products):")
	productCache.Print()

	// Access product (moves to front)
	price := productCache.Get(1001)
	fmt.Printf("Get product 1001 price: $%.2f\n", float64(price)/100)
	productCache.Print()

	// Add new product (evicts LRU)
	productCache.Put(1004, 29999) // Monitor: $299.99
	fmt.Println("Add product 1004 (evicts 1002 - Mouse):")
	productCache.Print()
}

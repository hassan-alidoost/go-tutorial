package main

import (
	"fmt"
)

type GoodStack[T any] struct {
	Items []T
}

const goodStackSize uint = 5

func NewStack[T any]() GoodStack[T] {
	return GoodStack[T]{Items: make([]T, 0, goodStackSize)}
}

func (s *GoodStack[T]) Push(item T) {
	if s.isFull() {
		fmt.Println("stack overflow!")
		return
	}
	s.Items = append(s.Items, item)
}

func (s *GoodStack[T]) Pop() T {
	if s.isEmpty() {
		fmt.Println("stack is empty!")
		var zero T
		return zero
	}

	lastItemIndex := len(s.Items) - 1
	lastItem := s.Items[lastItemIndex]
	s.Items = s.Items[:lastItemIndex]
	return lastItem
}

func (s *GoodStack[T]) isEmpty() bool {
	return len(s.Items) == 0
}

func (s *GoodStack[T]) isFull() bool {
	return len(s.Items) == int(goodStackSize)
}

func main() {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)

	fmt.Printf("stack items: %v\n", stack.Items)

	popItem := stack.Pop()
	fmt.Printf("Pop item %v\n", popItem)

	popItem = stack.Pop()
	fmt.Printf("Pop item %v\n", popItem)

	popItem = stack.Pop()
	fmt.Printf("Pop item %v\n", popItem)

	popItem = stack.Pop()
	fmt.Printf("Pop item %v\n", popItem)

	popItem = stack.Pop()
	fmt.Printf("Pop item %v\n", popItem)

	popItem = stack.Pop()
	fmt.Printf("Pop item %v\n", popItem)
}

package main

import (
	"fmt"
)

type Stack struct {
	Items []any
}

const stackSize uint = 5

func Build() Stack {
	return Stack{Items: make([]any, 0, stackSize)}
}

func (s *Stack) Push(item any) {
	if s.isFull() {
		fmt.Println("stack overflow!")
		return
	}
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() any {
	if s.isEmpty() {
		fmt.Println("stack is empty!")
		return nil
	}

	lastItemIndex := len(s.Items) - 1
	lastItem := s.Items[lastItemIndex]
	s.Items = s.Items[:lastItemIndex]
	return lastItem
}

func (s *Stack) isEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack) isFull() bool {
	return len(s.Items) == int(stackSize)
}

func main() {

	stack := Build()

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

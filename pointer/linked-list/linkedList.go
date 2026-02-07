package main

import "fmt"

type ListNode struct {
	Value int
	Next *ListNode
}

func main() {

	node1 := ListNode{Value: 3}
	node2 := ListNode{Value: 2}
	node1.Next = &node2
	node3 := ListNode{Value: 0}
	node2.Next = &node3
	node4 := ListNode{Value: -4}
	node3.Next = &node4
	node4.Next = &node2

	fmt.Println(hasCycle(&node1))
}
//
//
//          <---------------- 
//          |               |
//          |               |   
//     3 -> 2 -> 0 -> -4 -> 6
//1:   f
//1:   s
//2:        s    f 
//3:             s          f
//4:             f     s
//5:                        s
//5:                        f
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
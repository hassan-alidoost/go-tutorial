package main

import "fmt"


func main() {

	total := sumNumbers([]int{1, 2, 3, 4, 5})
	fmt.Println(total)

	total = sumNumbers([]int{-1, 0, 1})
	fmt.Println(total)

	total = sumNumbers([]int{})
	fmt.Println(total)
}


func sumNumbers(slice []int) int {
	var total int

	for _, v := range slice {
		total += v
	}

	return total
}
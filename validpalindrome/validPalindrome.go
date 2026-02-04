package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(12345))

}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)

	var bakward string

	for i := len(str) - 1; i >= 0; i-- {
		bakward = bakward + string(str[i])
	}

	return bakward == str
}

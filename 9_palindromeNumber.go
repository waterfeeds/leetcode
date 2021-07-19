package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(-10))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	return x == reverseInteger(x)
}

func reverseInteger(x int) int {
	var reverseInt int
	for x >= 10 {
		if reverseInt == 0 {
			reverseInt += x%10
		} else {
			reverseInt = reverseInt * 10 + x%10
		}
		x = x/10
	}

	reverseInt = reverseInt * 10 + x
	return reverseInt
}


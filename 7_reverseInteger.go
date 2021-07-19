package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	if x == 0 {
		return x
	}

	var reverseInt int
	if x > 0 {
		reverseInt = reversePositive(x)
	} else {
		reverseInt = -1*reversePositive(-1*x)
	}

	if reverseInt < -1 << 31 || reverseInt > 1<<31 - 1 {
		return 0
	}

	return reverseInt
}

func reversePositive(x int) int {
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
package main

import "fmt"

func main() {
	fmt.Println(divide(10, 3))
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	var isPositive = true
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		isPositive = false
	}

	dividend = abs(dividend)
	divisor = abs(divisor)
	if dividend < divisor {
		return 0
	}

	var result int
	var step = 1
	for dividend >= divisor << step {
		step ++
	}

	for dividend >= divisor {
		for dividend < divisor << step {
			step --
		}
		dividend -= divisor << step
		result += leftmove(step)
	}

	if !isPositive {
		result = -result
	}

	if result < -1<<31 || result > 1<<31-1 {
		result = 1<<31-1
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func leftmove(step int) int {
	var result = 1
	for i := 0;i < step;i ++ {
		result += result
	}

	return result
}

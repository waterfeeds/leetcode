package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var fn = make([]int, n+1)
	fn[1] = 1
	fn[2] = 2
	var index = 3
	for index <= n {
		fn[index] = fn[index-1] + fn[index-2]
		index++
	}

	return fn[n]
}

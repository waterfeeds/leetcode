package main

func numTrees(n int) int {
	var fn = make([]int, n+1)
	fn[0] = 1
	fn[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			fn[i] += fn[j]*fn[i-j-1]
		}
	}

	return fn[n]
}

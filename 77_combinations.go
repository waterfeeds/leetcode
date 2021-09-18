package main

func combine(n int, k int) [][]int {
	var result = make([][]int, 0)
	backtraceCombine(n, 1, k, []int{}, &result)
	return result
}

func backtraceCombine(n int, start int, k int, solution []int, result *[][]int) {
	if len(solution) == k {
		*result = append(*result, copylist(solution))
		return
	}

	if n-start+1+len(solution) < k {
		return
	}

	for index := start; index <= n; index ++ {
		solution = append(solution, index)
		backtraceCombine(n, index+1, k, solution, result)
		solution = solution[:len(solution)-1]
	}
}


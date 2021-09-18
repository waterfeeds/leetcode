package main

func subsets(nums []int) [][]int {
	var result = make([][]int, 0)
	backtraceSubsets(nums, 0, []int{}, &result)
	return result
}

func backtraceSubsets(nums []int, start int, solution []int, result *[][]int) {
	var length = len(nums)
	if start <= length {
		*result = append(*result, copylist(solution))
	}

	for index := start; index < length; index++ {
		solution = append(solution, nums[index])
		backtraceSubsets(nums, index+1, solution, result)
		solution = solution[:len(solution)-1]
	}
}

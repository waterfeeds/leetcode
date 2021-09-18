package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	var result = make([][]int, 0)
	backtraceSubsetsWithDup(nums, 0, []int{}, &result)
	return result
}

func backtraceSubsetsWithDup(nums []int, start int, solution []int, result *[][]int) {
	if start <= len(nums) {
		*result = append(*result, copylist(solution))
	}

	for index := start; index < len(nums); index++ {
		if index > start && nums[index] == nums[index-1] {
			continue
		}
		solution = append(solution, nums[index])
		backtraceSubsetsWithDup(nums, index+1, solution, result)
		solution = solution[:len(solution)-1]
	}
}

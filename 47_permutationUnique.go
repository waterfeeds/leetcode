package main

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var result = make([][]int, 0)

	backtracePermutationUnique(nums, 0, &result)
	return result
}

func backtracePermutationUnique(nums []int, start int, result *[][]int) {
	if start == len(nums) - 1 {
		*result = append(*result, copylist(nums))
		return
	}

	var uniqueMap = make(map[int]int)
	for i := start; i < len(nums); i++ {
		var num = nums[i]
		if _, ok := uniqueMap[num]; ok {
			continue
		}
		uniqueMap[num] = i

		swap(nums, i, start)
		backtracePermutationUnique(nums, start+1, result)
		swap(nums, i, start)
	}
}

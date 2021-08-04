package main

func permute(nums []int) [][]int {
	var result = make([][]int, 0)
	backtracePermutation(nums, 0, &result)

	return result
}

func backtracePermutation(nums []int, start int, result *[][]int) {
	if start == len(nums)-1 {
		*result = append(*result, copylist(nums))
		return
	}

	for i := start; i < len(nums); i++ {
		swap(nums, start, i)
		backtracePermutation(nums, start+1, result)
		swap(nums, start, i)
	}
}

func swap(nums []int, i, j int) {
	var temp = nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

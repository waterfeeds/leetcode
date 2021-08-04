package main

func jump2(nums []int) int {
	var step int
	var maxPosition int
	var end int

	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			step++
		}
	}

	return step
}

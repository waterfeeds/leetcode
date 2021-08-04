package main

func nextPermutation(nums []int)  {
	if len(nums) <= 1 {
		return
	}

	var length = len(nums)

	var i int
	for i = length-2; i >= 0; i--{
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i < 0 {
		nums = reverse31(nums, length)
		return
	}

	var nearLargeIdx = i+1
	for j := i+2;j <= length-1;j++ {
		if nums[j] > nums[i] && nums[j] <= nums[nearLargeIdx] {
			nearLargeIdx = j
		}
	}

	var tmp = nums[i]
	nums[i] = nums[nearLargeIdx]
	nums[nearLargeIdx] = tmp

	for k := 0; i+1+k < length-1-k;k ++ {
		var tmp = nums[i+1+k]
		nums[i+1+k] = nums[length-1-k]
		nums[length-1-k] = tmp
	}
}

func reverse31(nums []int, length int) []int {
	for i := 0;i < length-1-i;i ++ {
		var tmp = nums[i]
		nums[i] = nums[length-1-i]
		nums[length-1-i] = tmp
	}
	return nums
}

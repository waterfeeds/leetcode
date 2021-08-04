package main

func firstMissingPositive(nums []int) int {
	var length = len(nums)
	var temp int
	for i := 0; i < length; i++ {
		if nums[i] < 1 || nums[i] > length {
			continue
		}
		for nums[i] >= 1 && nums[i] <= length && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			temp = nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = temp
		}
	}

	var missingPositive = length+1
	for j := 0; j < length; j++ {
		if nums[j] != j+1 {
			missingPositive = j+1
			break
		}
	}

	return missingPositive
}

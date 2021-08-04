package main

func sortColors(nums []int) {
	var length = len(nums)
	var ptr = 0
	var temp int
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			if nums[ptr] > 0 {
				temp = nums[i]
				nums[i] = nums[ptr]
				nums[ptr] = temp
			}
			ptr++
		}
	}

	for j := ptr; j < length; j++ {
		if nums[j] == 1 {
			if nums[ptr] > 1 {
				temp = nums[j]
				nums[j] = nums[ptr]
				nums[ptr] = temp
			}
			ptr++
		}
	}
}

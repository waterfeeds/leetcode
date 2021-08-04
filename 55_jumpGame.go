package main

func canJump(nums []int) bool {
	var farPosition int
	for i := 0; i < len(nums); i++ {
		if i <= farPosition && farPosition < i+nums[i] {
			farPosition = i+nums[i]
		}
		if farPosition >= len(nums)-1 {
			return true
		}
	}

	return false
}

package main

func trap(height []int) int {
	var length = len(height)
	if length <= 1 {
		return 0
	}

	var left, right = 0, length-1
	var leftMax, rightMax int

	var trapRain int

	for left < right {
		leftVal, rightVal := height[left], height[right]
		leftMax = max(leftMax, leftVal)
		rightMax = max(rightMax, rightVal)

		if leftVal <= rightVal {
			trapRain += leftMax - leftVal
			left++
		} else {
			trapRain += rightMax- rightVal
			right--
		}
	}

	return trapRain
}

package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}

func maxArea(height []int) int {
	var left, right = 0, len(height)-1
	var area = 0
	var currentArea = 0

	for left < right {
		moveLeft := height[left] < height[right]
		if moveLeft {
			currentArea = (right-left)*height[left]
			left++
		} else {
			currentArea = (right-left)*height[right]
			right--
		}

		if currentArea > area {
			area = currentArea
		}
	}

	return area
}
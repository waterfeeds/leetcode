package main

func largestRectangleArea(heights []int) int {
	var length = len(heights)
	var stack = make([]int, 0)
	var left = make([]int, length)
	var right = make([]int, length)

	for i := 0; i < length; i++ {
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	stack = []int{}
	for i := length-1; i >= 0; i-- {
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = length
		} else {
			right[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	var area = 0
	for i := 0; i < length; i++ {
		area = max(area, heights[i]*(right[i]-left[i]-1))
	}
	return area
}

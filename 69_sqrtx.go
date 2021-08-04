package main

func mySqrt(x int) int {
	var left, right = 1, x
	for left <= right {
		mid := (left+right)/2
		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			left = mid+1
		} else {
			right = mid-1
		}
	}

	return right
}

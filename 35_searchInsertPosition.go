package main

func searchInsert(nums []int, target int) int {
	var length = len(nums)
	if length < 1 {
		return 0
	}

	if target < nums[0] {
		return 0
	}
	if target > nums[length-1] {
		return length
	}

	var left, right = 0, length-1
	for left < right {
		var mid = (left+right)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid+1
		} else {
			right = mid-1
		}
	}

	if nums[left] >= target {
		return left
	}

	return left+1
}

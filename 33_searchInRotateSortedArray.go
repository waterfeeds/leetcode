package main

func search(nums []int, target int) int {
	var length = len(nums)
	if length < 1 || length == 1 && nums[0] != target {
		return -1
	}

	if length == 1 {
		return 0
	}

	var reverseIdx int
	for i := 0; i < length-1; i++ {
		if nums[i] > nums[i+1] {
			reverseIdx = i
			break
		}
	}

	if reverseIdx == length-1 {
		return searchBinary(nums, 0, length-1, target)
	}

	if searchIdx := searchBinary(nums, 0, reverseIdx, target); searchIdx >= 0 {
		return searchIdx
	}

	return searchBinary(nums, reverseIdx+1, length-1, target)
}

func searchBinary(nums []int, left, right, target int) int {
	if target < nums[left] || target > nums[right] {
		return -1
	}

	for left <= right {
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

	return -1
}

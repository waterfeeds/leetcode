package main

func searchRotateArray(nums []int, target int) bool {
	var length = len(nums)
	var left, right = 0, length-1
	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			return true
		}

		if nums[mid] == nums[left] && nums[mid] == nums[right] {
			left++
			right--
			continue
		}
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

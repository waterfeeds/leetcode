package main

func searchRange(nums []int, target int) []int {
	var result = make([]int, 0)
	var left, right = 0, len(nums)-1

	for left <= right {
		var mid = (left+right)/2
		if nums[mid] == target {
			var firstIdx, lastIdx = mid, mid
			for firstIdx >= 0 {
				if nums[firstIdx] != target {
					break
				}
				firstIdx --
			}
			for lastIdx < len(nums) {
				if nums[lastIdx] != target {
					break
				}
				lastIdx ++
			}
			result = append(result, firstIdx+1, lastIdx-1)
			return result
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	result = append(result, -1, -1)
	return result
}


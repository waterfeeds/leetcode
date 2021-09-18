package main

func removeDuplicates2(nums []int) int {
	return removeDuplicatesK(nums, 2)
}

func removeDuplicatesK(nums []int, k int) int {
	var length = len(nums)
	if length <= k {
		return length
	}

	var fast, slow = k, k
	for fast < length {
		if nums[fast] != nums[slow-k] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}

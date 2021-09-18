package main

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int)  {
	var mergeI = m+n-1
	var i, j = m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[mergeI] = nums1[i]
			i--
		} else {
			nums1[mergeI] = nums2[j]
			j--
		}
		mergeI--
	}

	for j >= 0 {
		nums1[mergeI] = nums2[j]
		j--
		mergeI--
	}
}

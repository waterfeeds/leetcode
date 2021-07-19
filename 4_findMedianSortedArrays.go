package main

import "fmt"

func main() {
	nums1 := []int{0,0}
	nums2 := []int{0,0}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)

	if length % 2 != 0 {
		return findOddMedianSortedArrays(nums1, nums2)
	}
	return findEvenMedianSortedArrays(nums1, nums2)
}

func findOddMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j = 0, 0
	var len1 = len(nums1)
	var len2 = len(nums2)
	var iterIdx = 0
	var medIdx = (len1+len2)/2

	for i<len1 && j<len2 {
		if iterIdx == medIdx {
			if nums1[i] <= nums2[j] {
				return float64(nums1[i])
			} else {
				return float64(nums2[j])
			}
		}
		iterIdx ++
		if nums1[i] <= nums2[j] {
			i++
		} else {
			j++
		}
	}

	for i<len1 {
		if iterIdx == medIdx {
			return float64(nums1[i])
		}
		iterIdx++
		i++
	}

	for j<len2 {
		if iterIdx == medIdx {
			return float64(nums2[j])
		}
		iterIdx++
		j++
	}

	return 0
}

func findEvenMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j = 0, 0
	var len1 = len(nums1)
	var len2 = len(nums2)
	var iterIdx = 0
	var medIdx = (len1+len2)/2-1
	var medVal = 0
	var nextMedIdx = medIdx+1
	var nextMedVal = 0
	var nextMedGot = false

	for i<len1 && j<len2 {
		if iterIdx == medIdx {
			if nums1[i] <= nums2[j] {
				medVal = nums1[i]
			} else {
				medVal = nums2[j]
			}
		}
		if iterIdx == nextMedIdx {
			if nums1[i] <= nums2[j] {
				nextMedVal = nums1[i]
			} else {
				nextMedVal = nums2[j]
			}
			nextMedGot = true
		}

		if nextMedGot {
			return (float64(medVal) + float64(nextMedVal) ) / 2
		}
		iterIdx ++
		if nums1[i] <= nums2[j] {
			i++
		} else {
			j++
		}
	}

	for i<len1 {
		if iterIdx == medIdx {
			medVal = nums1[i]
		}
		if iterIdx == nextMedIdx {
			nextMedVal = nums1[i]
			nextMedGot = true
		}
		if nextMedGot {
			return (float64(medVal) + float64(nextMedVal) ) / 2
		}
		iterIdx++
		i++
	}

	for j<len2 {
		if iterIdx == medIdx {
			medVal = nums2[j]
		}
		if iterIdx == nextMedIdx {
			nextMedVal = nums2[j]
			nextMedGot = true
		}
		if nextMedGot {
			return (float64(medVal) + float64(nextMedVal) ) / 2
		}
		iterIdx++
		j++
	}

	return 0
}
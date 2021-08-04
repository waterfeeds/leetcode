package main

func searchMatrix(matrix [][]int, target int) bool {
	var m, n = len(matrix), len(matrix[0])

	var left, right = 0, m*n-1
	for left <= right {
		mid := (left+right)/2
		if matrix[mid/n][mid%n] == target {
			return true
		}
		if matrix[mid/n][mid%n] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

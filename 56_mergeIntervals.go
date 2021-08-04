package main

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	return mergeDfs(intervals, 0, len(intervals)-1)
}

func mergeDfs(intervals [][]int, left, right int) [][]int {
	if left == right {
		return [][]int{intervals[left]}
	}

	var mid = (left+right)/2
	leftIntervals := mergeDfs(intervals, left, mid)
	rightIntervals := mergeDfs(intervals, mid+1, right)

	return mergeTwo(leftIntervals, rightIntervals)
}

func mergeTwo(left, right [][]int) [][]int {
	var i, j = 0, 0
	var m, n = len(left), len(right)
	var result = make([][]int, 0)
	var merged []int
	for i < m && j < n {
		var current []int
		if left[i][1] < right[j][0] {
			current = left[i]
			i++
		} else if left[i][0] > right[j][1] {
			current = right[j]
			j++
		} else {
			current = []int{min(left[i][0], right[j][0]), max(left[i][1], right[j][1])}
			i++
			j++
		}

		if merged == nil {
			merged = current
			continue
		}

		if merged[1] < current[0] {
			result = append(result, []int{merged[0], merged[1]})
			merged = current
		} else {
			merged = []int{min(current[0], merged[0]), max(current[1], merged[1])}
		}
	}

	for i < m {
		var current = left[i]
		i++

		if merged[1] < current[0] {
			result = append(result, []int{merged[0], merged[1]})
			merged = current
		} else {
			merged = []int{min(current[0], merged[0]), max(current[1], merged[1])}
		}
	}

	for j < n {
		var current = right[j]
		j++

		if merged[1] < current[0] {
			result = append(result, []int{merged[0], merged[1]})
			merged = current
		} else {
			merged = []int{min(current[0], merged[0]), max(current[1], merged[1])}
		}
	}

	result = append(result, merged)
	return result
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

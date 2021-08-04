package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var current = newInterval
	var mergedIntervals = make([][]int, 0)
	var index int
	for index < len(intervals) {
		if current[1] < intervals[index][0] {
			mergedIntervals = append(mergedIntervals, current)
			return append(mergedIntervals, intervals[index:]...)
		}

		if current[0] > intervals[index][1] {
			mergedIntervals = append(mergedIntervals, intervals[index])
			index++
			continue
		}

		current = []int{
			min(current[0], intervals[index][0]),
			max(current[1], intervals[index][1]),
		}
		index++
	}

	return append(mergedIntervals, current)
}

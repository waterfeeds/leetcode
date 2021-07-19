package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{1,2,4,8,16,32,64,128}, 82))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var mostClosest = math.MinInt32
	for index, num := range nums {
		mostClosest = twoSumClosest(nums[index+1:], num, target, mostClosest)
		if mostClosest == target {
			break
		}
	}

	return mostClosest
}

func twoSumClosest(nums []int, num, target, mostClosest int) int {
	i, j := 0, len(nums)-1
	for i < j {
		currentClosest := nums[i] + nums[j] + num
		if currentClosest == target {
			return target
		}
		if currentClosest < target {
			i ++
		}
		if currentClosest > target {
			j--
		}
		if math.Abs(float64(currentClosest - target)) < math.Abs(float64(mostClosest - target)) {
			mostClosest = currentClosest
		}
	}

	return mostClosest
}

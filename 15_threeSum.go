package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	results := make([][]int, 0)
	for index := range nums {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		result := twoSumTemp(nums[index+1:], -nums[index])
		for _, r := range result {
			if len(results) < 1 || !eqInts(r, results[len(results)-1]){
				results = append(results, r)
			}
		}
	}

	return results
}

func twoSumTemp(nums []int, target int) [][]int {
	results := make([][]int, 0)
	var sMap = make(map[int]int)
	for index, num := range nums {
		if _, ok := sMap[target-num]; ok {
			result := []int{-target, target-num, num}
			if len(results) < 1 {
				results = append(results, result)
			} else {
				lastResult := results[len(results)-1]
				if compareInts(result[1:], lastResult[1:]) > 0 {
					results = append(results, result)
				} else if compareInts(result[1:], lastResult[1:]) < 0 {
					results = append([][]int{result}, results...)
				}
			}
		} else {
			sMap[num] = index
		}
	}

	return results
}

func eqInts(xInts, yInts []int) bool {
	if xInts[2] != yInts[2] || xInts[1] != yInts[1] || xInts[0] != yInts[0] {
		return false
	}

	return true
}

func compareInts(xInts, yInts []int) int {
	for index := range xInts {
		if xInts[index] > yInts[index] {
			return 1
		}
		if xInts[index] < yInts[index] {
			return -1
		}
	}

	return 0
}
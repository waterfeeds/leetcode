package main

import "fmt"

func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	rList := make([]int, 0)
	fMap := make(map[int]int)

	for index, num := range nums {
		if fIdx, ok := fMap[target-num]; ok {
			rList = append(rList, fIdx, index)
			break
		}
		if _, ok := fMap[num]; !ok {
			fMap[num] = index
		}
	}

	return rList
}

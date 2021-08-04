package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var result = [][]int{}
	backtrace(candidates, 0, []int{}, target, &result)

	return result
}

func backtrace(candidates []int, curIndex int, solution []int, target int, result *[][]int){
	if target == 0 {
		*result = append(*result, copylist(solution))
		return
	}

	if target < candidates[curIndex] {
		return
	}

	for index := curIndex; index < len(candidates); index ++ {
		solution = append(solution, candidates[index])
		backtrace(candidates, index, solution, target-candidates[index], result)
		solution = solution[:len(solution)-1]
	}
}

func copylist(list []int) []int {
	var cList = make([]int, 0, len(list))
	for _, element := range list {
		cList = append(cList, element)
	}

	return cList
}

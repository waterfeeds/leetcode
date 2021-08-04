package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var result = [][]int{}
	backtraceSum2(candidates, 0, []int{}, target, &result)

	return result
}

func backtraceSum2(candidates []int, start int, solution []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, copylist(solution))
		return
	}

	if start >= len(candidates) || target < candidates[start] {
		return
	}

	for index := start; index < len(candidates); index++ {
		if index > start && candidates[index] == candidates[index-1] {
			continue
		}
		solution = append(solution, candidates[index])
		backtraceSum2(candidates, index+1, solution, target-candidates[index], result)
		solution = solution[:len(solution)-1]
	}
}

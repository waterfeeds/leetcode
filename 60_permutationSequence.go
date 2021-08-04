package main

import "fmt"

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}

	var bitArr = make([]int, 0)

	var permutation = 1
	var permutationIndex = 1

	for permutationIndex < n {
		bitArr = append(bitArr, permutation)
		permutationIndex++
		permutation *= permutationIndex
	}

	var result string
	var currentK = k
	var permutationMap = make(map[int]int)

	for i := len(bitArr)-1; i >= 0; i-- {
		var num, j int
		for j = 1; j <= n; j++ {
			if _, ok := permutationMap[j]; !ok {
				num ++
				if num >= (currentK-1)/bitArr[i]+1 {
					break
				}
			}
		}

		result += fmt.Sprintf("%d", j)
		permutationMap[j] = j
		currentK = currentK%bitArr[i]

		if currentK == 1 {
			for j := 1; j <= n; j++ {
				if _, ok := permutationMap[j]; !ok {
					result += fmt.Sprintf("%d", j)
				}
			}
			return result
		}

		if currentK == 0 {
			for j := n; j >= 1; j-- {
				if _, ok := permutationMap[j]; !ok {
					result += fmt.Sprintf("%d", j)
				}
			}
			return result
		}
	}

	return result
}

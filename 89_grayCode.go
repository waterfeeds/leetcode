package main

func grayCode(n int) []int {
	var head = 1
	var result = []int{0}

	for i := 1; i <= n; i++ {
		for j := len(result)-1; j >= 0; j-- {
			result = append(result, head+result[j])
		}
		head *= 2
	}

	return result
}

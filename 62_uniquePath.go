package main

func uniquePaths(m int, n int) int {
	var fn = make([][]int, m+1)
	for rowId := range fn {
		fn[rowId] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		fn[i][1] = 1
	}
	for j := 1; j <= n; j++ {
		fn[1][j] = 1
	}

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			fn[i][j] = fn[i-1][j] + fn[i][j-1]
		}
	}

	return fn[m][n]
}

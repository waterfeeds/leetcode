package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	var m, n, t = len(s1), len(s2), len(s3)
	if m+n != t {
		return false
	}

	var fn = make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		fn[i] = make([]bool, n+1)
	}

	fn[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			p := i+j-1
			if i > 0 {
				fn[i][j] = fn[i][j] || (fn[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				fn[i][j] = fn[i][j] || (fn[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}

	return fn[m][n]
}

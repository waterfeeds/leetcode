package main

func numDecodings(s string) int {
	var length = len(s)
	if length < 1 || s[0] == '0' {
		return 0
	}
	var fn = make([]int, length+1)
	fn[0] = 1
	fn[1] = 1

	for i := 2; i <= length; i++ {
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] >= '0' && s[i-1] <= '6') {
			fn[i] += fn[i-2]
		}
		if s[i-1] >= '1' && s[i-1] <= '9' {
			fn[i] += fn[i-1]
		}
		if fn[i] == 0 {
			return 0
		}
	}

	return fn[length]
}

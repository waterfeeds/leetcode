package main

func longestValidParentheses(s string) int {
	var left, right, maxLength = 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left ++
		} else {
			right ++
		}

		if left == right {
			maxLength = max(left*2, maxLength)
		}

		if left < right {
			left = 0
			right = 0
		}
	}
	left = 0
	right = 0

	for i := len(s)-1; i >= 0;i-- {
		if s[i] == '(' {
			left ++
		} else {
			right ++
		}
		if left == right {
			maxLength = max(left*2, maxLength)
		}

		if left > right {
			left = 0
			right = 0
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

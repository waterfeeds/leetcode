package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("bb"))
}

func longestPalindrome(s string) string {
	var cIdx = 0
	length := len(s)
	var maxLen int
	var maxSubstring string

	for cIdx < length {
		if len(maxSubstring) == 0 {
			maxLen = 1
			maxSubstring = s[cIdx:cIdx+1]
		}
		for j := 1;cIdx-j >= 0 && cIdx+j < length;j ++ {
			if s[cIdx-j] == s[cIdx+j] {
				if 2*j+1 > maxLen {
					maxLen = 2*j+1
					maxSubstring = s[cIdx-j:cIdx+j+1]
				}
			} else {
				break
			}
		}
		for j := 1;cIdx-j+1 >= 0 && cIdx+j < length;j ++ {
			if s[cIdx-j+1] == s[cIdx+j] {
				if 2*j > maxLen {
					maxLen = 2*j
					maxSubstring = s[cIdx-j+1:cIdx+j+1]
				}
			} else {
				break
			}
		}

		cIdx++
	}

	return maxSubstring
}
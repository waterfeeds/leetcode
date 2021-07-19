package main

import "fmt"

func main() {
	s := "ab"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	lastIdxMap := make(map[uint8]int16)
	var maxLen int16 = 0
	var left int16 = 0
	var i int16 = 0
	for i = 0;int(i) < len(s);i++ {
		if left <= lastIdxMap[s[i]] {
			left = lastIdxMap[s[i]]
		}

		lastIdxMap[s[i]] = i+1

		if maxLen < i-left+1 {
			maxLen = i-left+1
		}
	}

	return int(maxLen)
}

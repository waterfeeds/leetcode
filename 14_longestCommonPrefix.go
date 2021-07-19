package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog","flow","flight"}))
}

func longestCommonPrefix(strs []string) string {
	var index int
	var commonPrefix string
	for {
		currentS := ""
		for _, str := range strs {
			if index >= len(str) {
				return commonPrefix
			}
			if len(currentS) < 1 {
				currentS = str[index:index+1]
			} else if currentS != str[index:index+1] {
				return commonPrefix
			}
		}
		commonPrefix += currentS
		index++
	}

	return commonPrefix
}

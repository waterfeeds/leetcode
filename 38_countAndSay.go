package main

import "fmt"

func countAndSay(n int) string {
	var curSay = "1"

	var index = 1
	for index < n {
		curSay = sayNext(curSay)
		index++
	}

	return curSay
}

func sayNext(curSay string) string {
	var nextSay string
	var curChar string
	var curCharCount int
	for i := 0; i < len(curSay); i++ {
		if curChar == "" {
			curChar = curSay[i:i+1]
			curCharCount = 1
		} else {
			if curSay[i:i+1] == curChar {
				curCharCount ++
			} else {
				nextSay += fmt.Sprintf("%d%s", curCharCount, curChar)
				curChar = curSay[i:i+1]
				curCharCount = 1
			}
		}
	}

	nextSay += fmt.Sprintf("%d%s", curCharCount, curChar)
	return nextSay
}

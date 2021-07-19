package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "aa"))
}

func isMatch(s string, p string) bool {
	if len(s) < 1 || len(p) < 1 {
		return false
	}

	var i, j = 0, 0
	for i < len(s) && j < len(p) {
		if s[i] != p[j] {
			return false
		}
		i++
		j++
	}

	if i < len(s) || j < len(p) {
		return false
	}
	return true
}

package main

import "strings"

func simplifyPath(path string) string {
	var stack = make([]string, 0)

	var length = len(path)
	var index = 0
	for index < length {
		if path[index] == '/' {
			for index < length && path[index] == '/'{
				index++
			}
			continue
		}

		startIndex := index
		for index < length && path[index] != '/' {
			index++
		}
		currentN := index-startIndex
		currentS := path[startIndex: index]
		if currentN == 1 {
			if currentS != "." {
				stack = append(stack, currentS)
			}
		} else if currentN == 2 && currentS == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, currentS)
		}
	}

	return "/" + strings.Join(stack, "/")
}

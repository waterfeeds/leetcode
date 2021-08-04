package main

func lengthOfLastWord(s string) int {
	var length = len(s)
	var index = length-1
	for index >= 0 {
		if s[index] != ' ' {
			break
		}
		index--
	}

	if index < 0 {
		return 0
	}

	var right = index
	index--
	for index >= 0 {
		if s[index] == ' ' {
			break
		}
		index--
	}

	return right-index
}

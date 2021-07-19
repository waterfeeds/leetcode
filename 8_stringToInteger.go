package main

import (
	"fmt"
)

func main() {
	fmt.Println(myAtoi("  -42"))
}

func myAtoi(s string) int {
	noBlankS := readNoBlank(s)
	if len(noBlankS) < 1 {
		return 0
	}

	isPositive, numberS := readPositive(noBlankS)

	numberInt := readNumber(isPositive, numberS)

	return numberInt
}

func readNoBlank(s string) string {
	for i := 0; i < len(s); i ++ {
		if s[i:i+1] == " " {
			continue
		}
		return s[i:]
	}

	return ""
}

func readPositive(s string) (isPositive bool, numberS string) {
	if s[0:1] == "-" {
		return false, s[1:]
	} else if s[0:1] == "+" {
		return true, s[1:]
	}

	return true, s
}

func readNumber(isPositive bool, s string) int {
	if len(s) < 1 {
		return 0
	}
	if s[0] < '0' || s[0] > '9' {
		return 0
	}

	var number = int(s[0]) - 48
	for i := 1; i < len(s);i++ {
		if s[i] >= '0' && s[i] <= '9' {
			number = number * 10 + int(s[i]) - 48
			if isPositive && number > 1<<31 - 1 {
				return 1<<31 - 1
			}
			if !isPositive && number > 1<<31 {
				return -1 << 31
			}
		} else {
			if !isPositive {
				number *= -1
			}
			return number
		}
	}

	if !isPositive {
		number *= -1
	}
	return number
}

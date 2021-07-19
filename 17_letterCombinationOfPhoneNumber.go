package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("723"))
}

func letterCombinations(digits string) []string {
	var allLetterCombinations = &[]string{}
	if len(digits) < 1 {
		return *allLetterCombinations
	}
	if len(digits) == 1 {
		letters := digitToLetter(digits[0])
		*allLetterCombinations = append(*allLetterCombinations, letters...)
		return *allLetterCombinations
	}

	letters := digitToLetter(digits[0])
	for _, letter := range letters {
		letterCombinationsDFS(digits[1:], letter, allLetterCombinations)
	}

	return *allLetterCombinations
}

func letterCombinationsDFS(digits string, prefix string, allLetterCombinations *[]string) {
	if len(digits) < 1 {
		*allLetterCombinations = append(*allLetterCombinations, prefix)
		return
	}
	letters := digitToLetter(digits[0])
	for _, letter := range letters {
		letterCombinationsDFS(digits[1:], prefix+letter, allLetterCombinations)
	}
}

func digitToLetter(digit uint8) []string {
	if digit >= '2' && digit <= '6' {
		return digitToLetterWithThree(digit)
	}

	return digitToLetterWithFour(digit)
}

func digitToLetterWithThree(digit uint8) []string {
	return []string{string(rune((digit-49)*3+94)), string(rune((digit-49)*3+95)), string(rune((digit-49)*3+96))}
}

func digitToLetterWithFour(digit uint8) []string {
	switch digit {
	case '7':
		return []string{"p", "q", "r", "s"}
	case '8':
		return []string{"t", "u", "v"}
	case '9':
		return []string{"w", "x", "y", "z"}
	}

	return []string{}
}
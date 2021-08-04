package main

func plusOne(digits []int) []int {
	var highBit = 1
	for index := len(digits)-1; index >= 0; index -- {
		current := digits[index] + highBit
		digits[index] = current%10
		highBit = current/10
		if highBit == 0 {
			return digits
		}
	}

	if highBit > 0 {
		digits = append([]int{highBit}, digits...)
	}

	return digits
}

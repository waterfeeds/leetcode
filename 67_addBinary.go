package main

import "fmt"

func addBinary(a string, b string) string {
	var highBit byte = '0'
	var index = 0
	var alen, blen = len(a), len(b)
	var current byte
	var result string

	for index < alen && index < blen {
		current = a[alen-1-index] + b[blen-1-index] + highBit - '0'*3
		result = fmt.Sprintf("%d", current%2) + result
		highBit = current/2 + 48
		index++
	}

	for index < alen {
		current = a[alen-1-index] + highBit - '0'*2
		result = fmt.Sprintf("%d", current%2) + result
		highBit = current/2 + 48
		index++
	}

	for index < blen {
		current = b[blen-1-index] + highBit - '0'*2
		result = fmt.Sprintf("%d", current%2) + result
		highBit = current/2 + 48
		index++
	}

	if highBit == '1' {
		result = "1" + result
	}

	return result
}

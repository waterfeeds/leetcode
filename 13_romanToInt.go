package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var valueSymbols = map[string]int{
		"M": 1000,
		"CM": 900,
		"D": 500,
		"CD": 400,
		"C": 100,
		"XC": 90,
		"L": 50,
		"XL": 40,
		"X": 10,
		"IX": 9,
		"V": 5,
		"IV": 4,
		"I": 1,
	}

	var romanValue int
	for i := 0;i < len(s); {
		if i+1 < len(s) {
			if romanInt, ok := valueSymbols[s[i:i+2]]; ok {
				romanValue += romanInt
				i += 2
				continue
			}
		}
		if romanInt, ok := valueSymbols[s[i:i+1]]; ok {
			romanValue += romanInt
			i += 1
			continue
		}
	}

	return romanValue
}

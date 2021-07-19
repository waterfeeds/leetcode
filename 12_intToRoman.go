package main

import "fmt"

func main() {
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	var valueSymbols = []struct {
		value int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman string
	for _, valueSymbol := range valueSymbols {
		for num >= valueSymbol.value {
			roman += valueSymbol.symbol
			num -= valueSymbol.value
		}
		if num == 0 {
			break
		}
	}

	return roman
}

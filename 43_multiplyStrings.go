package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var len1, len2 = len(num1), len(num2)
	var bigInt = make([]int8, len1+len2)

	for i := len1-1; i >= 0; i-- {
		for j := len2-1; j >= 0; j-- {
			sum := bigInt[i+j+1] + int8(num1[i]-48) * int8(num2[j]-48)
			bigInt[i+j+1] = sum%10
			bigInt[i+j] += sum/10
		}
	}

	var multiString string
	for index, bit := range bigInt {
		if index == 0 && bit == 0 {
			continue
		}
		multiString += fmt.Sprintf("%d", bit)
	}

	return multiString
}

package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n > 0 {
		return myPositivePow(x, n)
	}

	return 1.0/myPositivePow(x, -n)
}

type pow struct {
	index  int
	indexValue  float64
}

func myPositivePow(x float64, n int) float64 {
	var arrPow = make([]pow, 0)
	var sum float64 = 1
	var result float64 = x

	var index = 1
	arrPow = append(arrPow, pow{index: 1, indexValue: x})
	for index <= n {
		result *= result
		index *= 2
		arrPow = append(arrPow, pow{index: index, indexValue: result})
	}

	for i := len(arrPow)-1; i >= 0; i-- {
		if n >= arrPow[i].index {
			sum *= arrPow[i].indexValue
			n -= arrPow[i].index
		}
		if n == 0 {
			break
		}
	}

	return sum
}
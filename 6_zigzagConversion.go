package main

import (
	"fmt"
)

func main() {
	fmt.Println(len(convert("A", 2)))
}

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) < numRows {
		return s
	}

	numCols := len(s)*(numRows*numRows)/(2*numRows-1)
	if 2*numRows-1 > len(s) {
		numCols = len(s)-numRows+1
	}
	fmt.Println(numCols)

	var rr = make([][]uint8, numRows)
	for i := 0;i < numRows;i++ {
		rr[i] = make([]uint8, numCols)
	}

	var i, j = 0,0
	var iStep = 1
	for idx := 0;idx < len(s); idx++ {
		rr[i][j] = s[idx]
		i += iStep
		if iStep < 0 {
			j++
		}
		if i == numRows - 1 || i == 0 {
			iStep = -iStep
		}
	}

	var convertS string
	var reduceColsGot bool
	for i := 0;i < numRows;i ++ {
		for j := 0;j < numCols; j++ {
			if rr[i][j] > 0 {
				convertS += string(rr[i][j])
				if !reduceColsGot && j+numRows-1 < numCols && rr[i][j+numRows-1] == 0 {
					numCols = j + numRows - 1
					reduceColsGot = true
				}
			}
		}
	}

	return convertS
}
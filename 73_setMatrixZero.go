package main

func setZeroes(matrix [][]int)  {
	var rowMap = make(map[int]int)
	var columnMap = make(map[int]int)

	var rowL, columnL = len(matrix), len(matrix[0])

	for i := 0; i < rowL; i++ {
		for j := 0; j < columnL; j++ {
			if matrix[i][j] == 0 {
				if _, ok := rowMap[i]; !ok {
					rowMap[i] = i
				}
				if _, ok := columnMap[j]; !ok {
					columnMap[j] = j
				}
			}
		}
	}

	for i := 0; i < rowL; i++ {
		for j := 0; j < columnL; j++ {
			if _, ok := rowMap[i]; ok {
				matrix[i][j] = 0
				continue
			}
			if _, ok := columnMap[j]; ok {
				matrix[i][j] = 0
			}
		}
	}
}

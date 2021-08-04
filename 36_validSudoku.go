package main

func isValidSudoku(board [][]byte) bool {
	var rowsMap = make([]map[byte]int8, 9)
	var columnsMap = make([]map[byte]int8, 9)
	var boxsMap = make([]map[byte]int8, 9)

	for i := 0; i < 9; i++ {
		rowsMap[i] = make(map[byte]int8)
		columnsMap[i] = make(map[byte]int8)
		boxsMap[i] = make(map[byte]int8)
	}

	var val byte
	var ok bool
	for i := 0;i < 9; i++ {
		for j := 0;j < 9; j++ {
			val = board[i][j]
			if val == '.' {
				continue
			}

			if _, ok = rowsMap[i][val]; ok {
				return false
			} else {
				rowsMap[i][val] = 0
			}

			if _, ok = columnsMap[j][val]; ok {
				return false
			} else {
				columnsMap[j][val] = 0
			}

			var boxIndex = i/3*3 + j/3
			if _, ok = boxsMap[boxIndex][val]; ok {
				return false
			} else {
				boxsMap[boxIndex][val] = 0
			}
		}
	}
	return true
}

package main

func exist(board [][]byte, word string) bool {
	var m, n = len(board), len(board[0])
	var foundWord bool

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				var temp = board[i][j]
				board[i][j] = '0'
				backtraceExist(board, i, j, 1, word, &foundWord)
				board[i][j] = temp
				if foundWord {
					return true
				}
			}
		}
	}

	return foundWord
}

func backtraceExist(board [][]byte, startI int, startJ int, wordStart int, word string, foundWord *bool) {
	if *foundWord {
		return
	}

	if wordStart == len(word) {
		*foundWord = true
		return
	}

	var m, n = len(board), len(board[0])

	if startI < 0 || startI >= m || startJ < 0 || startJ >= n {
		return
	}

	if startI-1 >= 0 && board[startI-1][startJ] == word[wordStart] {
		var temp = board[startI-1][startJ]
		board[startI-1][startJ] = '0'
		backtraceExist(board, startI-1, startJ, wordStart+1, word, foundWord)
		board[startI-1][startJ] = temp
	}

	if startI+1 < m && board[startI+1][startJ] == word[wordStart] {
		var temp = board[startI+1][startJ]
		board[startI+1][startJ] = '0'
		backtraceExist(board, startI+1, startJ, wordStart+1, word, foundWord)
		board[startI+1][startJ] = temp
	}

	if startJ-1 >= 0 && board[startI][startJ-1] == word[wordStart] {
		var temp = board[startI][startJ-1]
		board[startI][startJ-1] = '0'
		backtraceExist(board, startI, startJ-1, wordStart+1, word, foundWord)
		board[startI][startJ-1] = temp
	}

	if startJ+1 < n && board[startI][startJ+1] == word[wordStart] {
		var temp = board[startI][startJ+1]
		board[startI][startJ+1] = '0'
		backtraceExist(board, startI, startJ+1, wordStart+1, word, foundWord)
		board[startI][startJ+1] = temp
	}
}

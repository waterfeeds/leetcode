package main

func generateMatrix(n int) [][]int {
	var matrix = make([][]int, n)
	for rowId := range matrix {
		matrix[rowId] = make([]int, n)
	}

	var number = 1
	var i, j = 0, 0
	var direction = "right"

	for i < n && j < n {
		matrix[i][j] = number

		switch direction {
		case "right":
			j++
			if j >= n || matrix[i][j] > 0 {
				i++
				j--
				direction = "down"
			}
		case "down":
			i++
			if i >= n || matrix[i][j] > 0 {
				i--
				j--
				direction = "left"
			}
		case "left":
			j--
			if j < 0 || matrix[i][j] > 0 {
				i--
				j++
				direction = "up"
			}
		case "up":
			i--
			if i < 0 || matrix[i][j] > 0 {
				i++
				j++
				direction = "right"
			}
		}

		number++
		if number > n*n {
			break
		}
	}

	return matrix
}

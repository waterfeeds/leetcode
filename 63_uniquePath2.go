package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var m, n = len(obstacleGrid), len(obstacleGrid[0])

	var hasStone bool
	for i := 0; i < m; i++ {
		if hasStone {
			obstacleGrid[i][0] = 0
			continue
		}
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
			hasStone = true
			continue
		}
		if obstacleGrid[i][0] == 0 {
			obstacleGrid[i][0] = 1
		}
	}

	if obstacleGrid[0][0] == 1 {
		hasStone = false
	} else {
		hasStone = true
	}
	for j := 1; j < n; j++ {
		if hasStone {
			obstacleGrid[0][j] = 0
			continue
		}
		if obstacleGrid[0][j] == 1 {
			obstacleGrid[0][j] = 0
			hasStone = true
			continue
		}
		if obstacleGrid[0][j] == 0 {
			obstacleGrid[0][j] = 1
		}
	}

	var up, left int
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			up, left = 0, 0
			if obstacleGrid[i-1][j] > 0 {
				up = obstacleGrid[i-1][j]
			}

			if obstacleGrid[i][j-1] > 0 {
				left = obstacleGrid[i][j-1]
			}

			obstacleGrid[i][j] = up+left
		}
	}

	return obstacleGrid[m-1][n-1]
}

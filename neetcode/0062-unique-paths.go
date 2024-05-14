// tags: 2d-dp, math(todo), medium

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		matrix[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = matrix[i+1][j] + matrix[i][j+1]
			}
		}
	}

	return matrix[0][0]
}

func uniquePaths(m int, n int) int {
	rowVals := make([]int, n)

	for row := range rowVals {
		rowVals[row] = 1
	}

	for row := m - 2; row >= 0; row-- {
		nextRowVals := make([]int, n)
		for col := n - 1; col >= 0; col-- {
			var right, down int
			if col+1 < n {
				right = nextRowVals[col+1]
			}
			down = rowVals[col]
			nextRowVals[col] += right + down
		}
		rowVals = nextRowVals
	}
	return rowVals[0]
}
// tags: dp, math(todo)

// func uniquePaths(m int, n int) int {
//     grid := make([][]int, m)
//     for row := range grid {
//         grid[row] = make([]int, n)
//     }

//     grid[m-1][n-1] = 1

//     for row := m-1; row >= 0; row-- {
//         for col := n-1; col >= 0; col-- {
//             var right, down int
//             if col+1 < n {
//                 right = grid[row][col+1]
//             }
//             if row+1 < m {
//                 down = grid[row+1][col]
//             }
//             grid[row][col] += right + down
//         }
//     }
//     return grid[0][0]
// }

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
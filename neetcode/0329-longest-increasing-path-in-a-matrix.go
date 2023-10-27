// tags: star2, dp, topological-sort

// time complexity: O(m*n)
// space complexity: O(m*n)
func longestIncreasingPath(matrix [][]int) int {
	dp := make(map[[2]int]int)
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if val, ok := dp[[2]int{row, col}]; ok {
			return val
		}

		pathVal := 1
		for _, direct := range [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} {
			nextRow, nextCol := row+direct[0], col+direct[1]
			if nextRow >= 0 && nextRow < len(matrix) && nextCol >= 0 && nextCol < len(matrix[0]) &&
				matrix[row][col] < matrix[nextRow][nextCol] {
				pathVal = max(pathVal, 1+dfs(nextRow, nextCol))
			}
		}
		dp[[2]int{row, col}] = pathVal

		return pathVal
	}

	var res int
	for row := range matrix {
		for col := range matrix[row] {
			temp := dfs(row, col)
			res = max(res, temp)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
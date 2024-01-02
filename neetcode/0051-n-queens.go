// tags: backtracking, star3, hard

// time complexity: O(N!)
// space complexity: O(N)
// `N` is board size n*n
func solveNQueens(n int) [][]string {
	col := make(map[int]bool)
	posDiag := make(map[int]bool)
	negDiag := make(map[int]bool)

	var (
		backtrack func(rowIdx int)
		res       [][]string
	)
	board := make([][]string, n)
	for rowIdx := range board {
		board[rowIdx] = make([]string, n)
		for colIdx := range board[rowIdx] {
			board[rowIdx][colIdx] = "."
		}
	}
	backtrack = func(rowIdx int) {
		if rowIdx == n {
			forAppend := make([]string, 0, len(board))
			for i := range board {
				forAppend = append(forAppend, strings.Join(board[i], ""))
			}
			res = append(res, forAppend)
			return
		}

		for i := 0; i < n; i++ {
			if col[i] || posDiag[rowIdx+i] || negDiag[rowIdx-i] {
				continue
			}

			col[i] = true
			posDiag[rowIdx+i] = true
			negDiag[rowIdx-i] = true
			board[rowIdx][i] = "Q"

			backtrack(rowIdx + 1)

			col[i] = false
			posDiag[rowIdx+i] = false
			negDiag[rowIdx-i] = false
			board[rowIdx][i] = "."
		}
	}
	backtrack(0)

	return res
}
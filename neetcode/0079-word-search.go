// time complexity: O(m*n*4^s)
// space complexity: O(m*n)
func exist(board [][]byte, word string) bool {
	var res bool
	for row, line := range board {
		for col := range line {
			res = res || dfs(board, word, row, col, 0)
			if res {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, row, col, idx int) bool {
	if idx == len(word) {
		return true
	}
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return false
	}
	if board[row][col] == '*' {
		return false
	}
	if word[idx] != board[row][col] {
		return false
	}

	visited := board[row][col]
	board[row][col] = '*'

	res := dfs(board, word, row, col-1, idx+1) ||
		dfs(board, word, row-1, col, idx+1) ||
		dfs(board, word, row, col+1, idx+1) ||
		dfs(board, word, row+1, col, idx+1)

	board[row][col] = visited

	return res
}
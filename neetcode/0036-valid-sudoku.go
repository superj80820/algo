// tags: arrays&hashing

// time complexity: O(n^2)
// space complexity: O(n^2)
func isValidSudoku(board [][]byte) bool {
	checkRows := make(map[int]map[byte]bool)
	checkCols := make(map[int]map[byte]bool)
	checkSquare := make(map[[2]int]map[byte]bool)

	for row, line := range board {
		for col, val := range line {
			if val == '.' {
				continue
			}

			squareKey := [2]int{row / 3, col / 3}

			if checkRows[row][val] ||
				checkCols[col][val] ||
				checkSquare[[2]int{row / 3, col / 3}][val] {
				return false
			}

			if _, ok := checkRows[row]; !ok {
				checkRows[row] = make(map[byte]bool)
			}
			if _, ok := checkCols[col]; !ok {
				checkCols[col] = make(map[byte]bool)
			}
			if _, ok := checkSquare[squareKey]; !ok {
				checkSquare[squareKey] = make(map[byte]bool)
			}
			checkRows[row][val] = true
			checkCols[col][val] = true
			checkSquare[squareKey][val] = true
		}
	}

	return true
}
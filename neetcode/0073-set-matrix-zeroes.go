// tags: math&geometry, medium

// time complexity: O(m*n)
// space complexity: O(1)
func setZeroes(matrix [][]int) {
	var isRowZero bool
	for rowIdx := range matrix {
		for colIdx := range matrix[rowIdx] {
			if matrix[rowIdx][colIdx] == 0 {
				if rowIdx == 0 {
					isRowZero = true
				} else {
					matrix[rowIdx][0] = 0
				}
				matrix[0][colIdx] = 0
			}
		}
	}
	for rowIdx := 1; rowIdx < len(matrix); rowIdx++ {
		for colIdx := 1; colIdx < len(matrix[rowIdx]); colIdx++ {
			if matrix[rowIdx][0] == 0 || matrix[0][colIdx] == 0 {
				matrix[rowIdx][colIdx] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for rowIdx := 1; rowIdx < len(matrix); rowIdx++ {
			matrix[rowIdx][0] = 0
		}
	}
	if isRowZero {
		for colIdx := 0; colIdx < len(matrix[0]); colIdx++ {
			matrix[0][colIdx] = 0
		}
	}
}
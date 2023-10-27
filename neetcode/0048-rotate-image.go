// tags: math&geometry

// time complexity: O(m*n)
// space complexity: O(1)
func rotate(matrix [][]int) {
	l, r := 0, len(matrix[0])-1

	for l < r {
		for i := 0; i < r-l; i++ {
			t, b := l, r

			temp := matrix[t][l+i]

			matrix[t][l+i] = matrix[b-i][l]

			matrix[b-i][l] = matrix[b][r-i]

			matrix[b][r-i] = matrix[t+i][r]

			matrix[t+i][r] = temp
		}
		l++
		r--
	}
}
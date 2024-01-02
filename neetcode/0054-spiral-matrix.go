// tags: math&geometry, medium

// time complexity: O(m*n)
// space complexity: O(1)
func spiralOrder(matrix [][]int) []int {
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1

	var res []int
	for l <= r && t <= b {
		for idx := l; idx <= r; idx++ {
			res = append(res, matrix[t][idx])
		}
		t++
		for idx := t; idx <= b; idx++ {
			res = append(res, matrix[idx][r])
		}
		r--
		if !(l <= r && t <= b) {
			break
		}
		for idx := r; idx >= l; idx-- {
			res = append(res, matrix[b][idx])
		}
		b--
		for idx := b; idx >= t; idx-- {
			res = append(res, matrix[idx][l])
		}
		l++
	}
	return res
}
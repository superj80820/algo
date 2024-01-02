// tags: bit-manipulation, medium

// time complexity: O(1)
// space complexity: O(1)
func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}
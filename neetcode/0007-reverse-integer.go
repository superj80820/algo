// tags: bit-manipulation, PR, medium

import "math"

// time complexity: O(n)
// space complexity: O(1)
// n is x length
func reverse(x int) int {
	// math.MinInt32: -2147483648
	// math.MaxInt32: 02147483647

	var res int
	for x != 0 {
		digit := x % 10
		x = x / 10

		// if you system let int is 32-bit, you need use this check way for avoid overflow
		// e.g. input: 1000000032 will overflow when reverse
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && digit < math.MinInt32%10) ||
			res > math.MaxInt32/10 || (res == math.MaxInt32/10 && digit > math.MaxInt32%10) {
			return 0
		}

		res = res*10 + digit
	}
	return res
}
// tags: math&geometry

// time complexity: O(logn)
// space complexity: O(logn)
func myPow(x float64, n int) float64 {
	var helper func(x float64, n int) float64
	helper = func(x float64, n int) float64 {
		if x == 0 {
			return 0
		}
		if n == 0 {
			return 1
		}
		res := helper(x, n/2)
		res = res * res
		if n%2 != 0 {
			return x * res
		}
		return res
	}
	res := helper(x, n)
	if n >= 0 {
		return res
	}
	return 1 / res
}

// tags: bit-manipulation, easy

// time complexity: O(n)
// space complexity: O(n)
func countBits(n int) []int {
	dp := make([]int, n+1)
	curPow := 1
	for i := 1; i <= n; i++ {
		dp[i] = 1 + dp[i-curPow]
		if i+1 == curPow*2 {
			curPow *= 2
		}
	}
	return dp
}
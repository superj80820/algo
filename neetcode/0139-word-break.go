// tags: star2, dp

// time complexity: O(n^3)
// space complexity: O(n)
// `n` is length of s
// `m` is length of wordDict
func wordBreak(s string, wordDict []string) bool {
	dp := make(map[int]bool)
	dp[len(s)] = true
	for i := len(s) - 1; i >= 0; i-- {
		for _, word := range wordDict {
			if len(s)-i >= len(word) && s[i:i+len(word)] == word {
				dp[i] = dp[i+len(word)]
			}
			if dp[i] {
				break
			}
		}
	}
	return dp[0]
}
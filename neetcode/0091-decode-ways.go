// tags: 1d-dp, star3, medium

// time complexity: O(n)
// space complexity: O(n)
func numDecodings(s string) int {
	dp := make(map[int]int)
	var dfs func(idx int) int
	dfs = func(idx int) int {
		if idx == len(s) {
			return 1
		}
		if s[idx] == '0' {
			return 0
		}

		if val, ok := dp[idx]; ok {
			return val
		}
		var res int
		res += dfs(idx + 1)
		if idx+1 < len(s) &&
			(s[idx] == '1' || s[idx] == '2' && s[idx+1] >= '0' && s[idx+1] <= '6') {
			res += dfs(idx + 2)
		}
		dp[idx] = res
		return res
	}
	return dfs(0)
}

// time complexity: O(n)
// space complexity: O(n)
func numDecodings(s string) int {
	dp := make(map[int]int, len(s)+1)
	for idx := len(s); idx >= 0; idx-- {
		if idx == len(s) {
			dp[idx] = 1
			continue
		}
		if s[idx] == '0' {
			dp[idx] = 0
			continue
		}

		var res int
		res += dp[idx+1]
		if idx+1 < len(s) &&
			(s[idx] == '1' || s[idx] == '2' && s[idx+1] >= '0' && s[idx+1] <= '6') {
			res += dp[idx+2]
		}
		dp[idx] = res
	}
	return dp[0]
}

// time complexity: O(n)
// space complexity: O(1)
func numDecodings(s string) int {
	var first, second int
	for idx := len(s); idx >= 0; idx-- {
		if idx == len(s) {
			first = 1
			continue
		}
		if s[idx] == '0' {
			first, second = 0, first
			continue
		}

		var res int
		res += first
		if idx+1 < len(s) &&
			(s[idx] == '1' || s[idx] == '2' && s[idx+1] >= '0' && s[idx+1] <= '6') {
			res += second
		}
		first, second = res, first
	}
	return first
}
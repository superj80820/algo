// tags: 2d-dp, star3, dfs

// time complexity: O(m*n)
// space complexity: O(m*n)
// `m` is length of s1
// `n` is length of s2
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	cache := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == len(s1) && j == len(s2) {
			return true
		}
		if val, ok := cache[[2]int{i, j}]; ok {
			return val
		}

		if i < len(s1) && s1[i] == s3[i+j] && dfs(i+1, j) {
			return true
		}
		if j < len(s2) && s2[j] == s3[i+j] && dfs(i, j+1) {
			return true
		}
		cache[[2]int{i, j}] = false
		return false
	}

	return dfs(0, 0)
}

// time complexity: O(m*n)
// space complexity: O(m*n)
// `m` is length of s1
// `n` is length of s2
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for row := range dp {
		dp[row] = make([]bool, len(s2)+1)
	}

	dp[0][0] = true

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = s2[j-1] == s3[i+j-1] && dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = s1[i-1] == s3[i+j-1] && dp[i-1][j]
			} else {
				dp[i][j] = s2[j-1] == s3[i+j-1] && dp[i][j-1] || s1[i-1] == s3[i+j-1] && dp[i-1][j]
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
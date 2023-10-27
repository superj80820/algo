// tags: 2d-dp, star2

// time complexity: O(n^3)
// space complexity: O(n^2)
func maxCoins(nums []int) int {
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l > r {
			return 0
		}
		if dp[l][r] != 0 {
			return dp[l][r]
		}
		for i := l; i <= r; i++ {
			coin := nums[l-1] * nums[i] * nums[r+1]
			coin += dfs(l, i-1) + dfs(i+1, r)
			dp[l][r] = max(dp[l][r], coin)
		}
		return dp[l][r]
	}
	return dfs(1, len(nums)-2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
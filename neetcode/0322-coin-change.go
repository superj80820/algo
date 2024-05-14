// tags: 1d-dp, star3, medium

package neetcode

import "math"

// time complexity: O(n*c)
// space complexity: O(n)
// `n` is amount number
// `c` is length of coins
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func coinChange(coins []int, amount int) int {
	dp := make(map[int]int)
	var dfs func(amount int) int
	dfs = func(amount int) int {
		if amount < 0 {
			return -1
		}
		if amount == 0 {
			return 0
		}
		if val, exists := dp[amount]; exists {
			return val
		}

		minCoins := math.MaxInt32
		for _, coin := range coins {
			result := dfs(amount - coin)
			if result >= 0 {
				minCoins = min(minCoins, result+1)
			}
		}

		if minCoins == math.MaxInt32 {
			dp[amount] = -1
		} else {
			dp[amount] = minCoins
		}

		return dp[amount]
	}

	return dfs(amount)
}

// tags: 2d-dp, star3, dfs, medium, practice-count:3

// time complexity: O(m*n)
// space complexity: O(m*n)
// `m` is amount
// `n` is length of coins
func change(amount int, coins []int) int {
	cache := make(map[[2]int]int)
	var dfs func(idx, amount int) int
	dfs = func(idx, amount int) int {
		if val, ok := cache[[2]int{idx, amount}]; ok {
			return val
		}
		if amount == 0 {
			return 1
		}
		if amount < 0 {
			return 0
		}
		if idx >= len(coins) {
			return 0
		}
		nextAmount := amount - coins[idx]
		res := dfs(idx, nextAmount) + dfs(idx+1, amount)
		cache[[2]int{idx, amount}] = res
		return res
	}
	return dfs(0, amount)
}

// time complexity: O(m*n)
// space complexity: O(m*n)
// `m` is amount
// `n` is length of coins
func change(amount int, coins []int) int {
	dp := make([][]int, amount+1)
	for row := range dp {
		dp[row] = make([]int, len(coins))
		if row == 0 {
			for col := range dp[row] {
				dp[row][col] = 1
			}
		}
	}

	for row := 1; row <= amount; row++ {
		for col := range dp[row] {
			var up, left int
			if row-coins[col] >= 0 {
				up = dp[row-coins[col]][col]
			}
			if col-1 >= 0 {
				left = dp[row][col-1]
			}
			dp[row][col] = up + left
		}
	}

	return dp[amount][len(coins)-1]
}

// time complexity: O(m*n)
// space complexity: O(m)
// `m` is amount
// `n` is length of coins
func change(amount int, coins []int) int {
	dpColVals := make([]int, amount+1)
	for col := range coins {
		for row := range dpColVals {
			if row == 0 {
				dpColVals[row] = 1
				continue
			}
			var up, left int
			if row-coins[col] >= 0 {
				up = dpColVals[row-coins[col]]
			}
			left = dpColVals[row]
			dpColVals[row] = up + left
		}
	}

	return dpColVals[amount]
}
// tags: star3, dp, dfs, todo-space-optimize

// time complexity: O(n)
// space complexity: O(n)
func maxProfit(prices []int) int {
	cache := make(map[[2]interface{}]int)
	var dfs func(idx int, isBuying bool) int

	dfs = func(idx int, isBuying bool) int {
		if idx >= len(prices) {
			return 0
		}
		if val, ok := cache[[2]interface{}{idx, isBuying}]; ok {
			return val
		}

		if isBuying {
			cooldown := dfs(idx+1, true)
			buy := dfs(idx+1, false) - prices[idx]
			cache[[2]interface{}{idx, true}] = max(buy, cooldown)
		} else {
			cooldown := dfs(idx+1, false)
			sell := dfs(idx+2, true) + prices[idx]
			cache[[2]interface{}{idx, false}] = max(sell, cooldown)
		}

		return cache[[2]interface{}{idx, isBuying}]
	}

	return dfs(0, true)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
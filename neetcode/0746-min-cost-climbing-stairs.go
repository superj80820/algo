// tags: 1d-dp, star2, easy, practice-count:2

// time complexity: O(n)
// space complexity: O(1)
func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// time complexity: O(n)
// space complexity: O(n)
func minCostClimbingStairs(cost []int) int {
	cache := make(map[int]int, len(cost))
	dp(0, cost, cache)
	return min(cache[0], cache[1])
}

func dp(idx int, cost []int, cache map[int]int) int {
	if idx >= len(cost) {
		return 0
	}
	if val, ok := cache[idx]; ok {
		return val
	}
	res := cost[idx] + min(dp(idx+1, cost, cache), dp(idx+2, cost, cache))
	cache[idx] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

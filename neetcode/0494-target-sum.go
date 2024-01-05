// tags: 2d-dp, star3, dfs, medium

// time complexity: O(n*t)
// space complexity: O(n*t)
// `n` is length of nums
// `t` is sum of nums
func findTargetSumWays(nums []int, target int) int {
	var dfs func(idx, total int) int
	cache := make(map[[2]int]int)
	dfs = func(idx, total int) int {
		if val, ok := cache[[2]int{idx, total}]; ok {
			return val
		}
		if idx == len(nums) {
			if total == target {
				return 1
			}
			return 0
		}
		res := dfs(idx+1, total+nums[idx]) + dfs(idx+1, total-nums[idx])
		cache[[2]int{idx, total}] = res
		return res
	}
	return dfs(0, 0)
}
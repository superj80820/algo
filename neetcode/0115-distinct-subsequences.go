// tags: star3, dp, lcs

// time complexity: O(m*n)
// space complexity: O(m*n)
// `m` is length of s
// `n` is length of t
func numDistinct(s string, t string) int {
	cache := make(map[[2]int]int)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if val, ok := cache[[2]int{i, j}]; ok {
			return val
		}
		if j == len(t) {
			return 1
		}
		if i == len(s) {
			return 0
		}

		if s[i] == t[j] {
			cache[[2]int{i, j}] = dfs(i+1, j+1) + dfs(i+1, j)
		} else {
			cache[[2]int{i, j}] = dfs(i+1, j)
		}
		return cache[[2]int{i, j}]
	}

	return dfs(0, 0)
}

// time complexity: O(m*n)
// space complexity: O(m*n)
// `m` is length of s
// `n` is length of t
func numDistinct(s string, t string) int {
	cache := make([][]int, len(s)+1)
	for i := len(s); i >= 0; i-- {
		cache[i] = make([]int, len(t)+1)
		for j := len(t); j >= 0; j-- {
			cache[i][j] = -1
		}
	}

	for i := len(s); i >= 0; i-- {
		for j := len(t); j >= 0; j-- {
			if cache[i][j] != -1 {
				continue
			}
			if j == len(t) {
				cache[i][j] = 1
				continue
			}
			if i == len(s) {
				cache[i][j] = 0
				continue
			}

			if s[i] == t[j] {
				cache[i][j] = cache[i+1][j+1] + cache[i+1][j]
			} else {
				cache[i][j] = cache[i+1][j]
			}
		}
	}

	return cache[0][0]
}
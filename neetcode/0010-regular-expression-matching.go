// tags: 2d-dp, hard

// time complexity: O(m*n)
// space complexity: O(m*n)
// `m` is length of s
// `n` is length of p
func isMatch(s string, p string) bool {
	cache := make(map[[2]int]bool)

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i >= len(s) && j >= len(p) {
			return true
		} else if j >= len(p) {
			return false // TODO: york
		}

		match := i < len(s) && (s[i] == p[j] || p[j] == '.')
		cacheKey := [2]int{i, j}
		if j+1 < len(p) && p[j+1] == '*' { // star case
			cache[cacheKey] = dfs(i, j+2) || (match && dfs(i+1, j))
		} else if match { // not star and match case
			cache[cacheKey] = dfs(i+1, j+1)
		} else {
			cache[cacheKey] = false
		}

		return cache[cacheKey]
	}

	return dfs(0, 0)
}

// time complexity: O(m*n)
// space complexity: O(m*n)
// `m` is length of s
// `n` is length of p
func isMatch(s string, p string) bool {
	cache := make([][]bool, len(s)+1)
	for row := range cache {
		cache[row] = make([]bool, len(p)+1)
	}
	cache[len(s)][len(p)] = true
	for i := len(s); i >= 0; i-- {
		for j := len(p); j >= 0; j-- {
			if i >= len(s) && j >= len(p) {
				continue
			} else if j >= len(p) {
				continue
			}
			match := i < len(s) && (s[i] == p[j] || p[j] == '.')
			if j+1 < len(p) && p[j+1] == '*' { // star case
				cache[i][j] = cache[i][j+2] || (match && cache[i+1][j])
			} else if match { // not star and match case
				cache[i][j] = cache[i+1][j+1]
			} else {
				cache[i][j] = false
			}
		}
	}

	return cache[0][0]
}
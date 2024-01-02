// tags: sliding-window, star3, medium

// time complexity: O(n)
// space complexity: O(n)
func lengthOfLongestSubstring(s string) int {
	check := make(map[byte]bool)
	var l, r, res int
	for ; r < len(s); r++ {
		for check[s[r]] {
			delete(check, s[l])
			l++
		}
		check[s[r]] = true
		res = max(res, r-l+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
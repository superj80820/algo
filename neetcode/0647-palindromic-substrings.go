// tags: 1d-dp, star1, medium

// time complexity: O(n^2)
// space complexity: O(1)
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	var (
		res  int
		l, r int
	)
	check := func() {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	for i := range s {
		l, r = i, i+1
		check()
		l, r = i, i
		check()
	}
	return res
}

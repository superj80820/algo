// tags: 1d-dp

// time complexity: O(n^2)
// space complexity: O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	var (
		maxStrLen int
		res       string
		l, r      int
	)
	for i := range s {
		check := func() {
			for l >= 0 && r < len(s) && s[l] == s[r] {
				if r-l+1 > maxStrLen {
					res = s[l : r+1]
					maxStrLen = r - l + 1
				}
				l--
				r++
			}
		}
		l, r = i, i+1
		check()
		l, r = i, i
		check()
	}
	return res
}

// time complexity: O(n^3)
// space complexity: O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	var (
		maxStrLen int
		res       string
		l, r      int
	)
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			check := func() {
				for l >= 0 && r < len(s) && s[l] == s[r] {
					if r-l+1 > maxStrLen {
						res = s[l : r+1]
						maxStrLen = r - l + 1
					}
					l--
					r++
				}
			}
			l, r = j, j+1
			check()
			l, r = j, j
			check()
		}
	}
	return res
}
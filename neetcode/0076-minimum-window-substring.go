// tags: star1, sliding-window

import "math"

// time complexity: O(n)
// space complexity: O(n)
func minWindow(s string, t string) string {
	var (
		resIdx     [2]int
		l, r       int
		have, need int
	)
	resLen := math.MaxInt32
	countS, countT := make(map[byte]int), make(map[byte]int)
	for idx := range t {
		countT[t[idx]]++
	}
	need = len(countT)

	for r < len(s) {
		if countT[s[r]] != 0 {
			countS[s[r]]++
			if countS[s[r]] == countT[s[r]] {
				have++
			}
		}
		for have == need {
			if r-l+1 < resLen {
				resLen = r - l + 1
				resIdx[0], resIdx[1] = l, r
			}
			if countT[s[l]] != 0 {
				countS[s[l]]--
				if countS[s[l]] < countT[s[l]] {
					have--
				}
			}
			l++
		}
		r++
	}

	if resLen == math.MaxInt32 {
		return ""
	}
	return s[resIdx[0] : resIdx[1]+1]
}
// tags: sliding-window, medium

// time complexity: O(n)
// space complexity: O(n)
func checkInclusion(s1 string, s2 string) bool {
	shortStr, longStr := s1, s2
	if len(s1) > len(s2) {
		return false
	}

	var shortCheck, longCheck [26]int

	for idx := range shortStr {
		shortCheck[shortStr[idx]-'a']++
		longCheck[longStr[idx]-'a']++
	}

	var match int
	for i := 0; i < 26; i++ {
		if shortCheck[i] == longCheck[i] {
			match++
		}
	}

	for l, r := 0, len(shortStr); r < len(longStr); l, r = l+1, r+1 {
		if match == 26 {
			return true
		}

		idx := longStr[l] - 'a'
		longCheck[idx]--
		if shortCheck[idx] == longCheck[idx] {
			match++
		} else if shortCheck[idx] == longCheck[idx]+1 {
			match--
		}

		idx = longStr[r] - 'a'
		longCheck[idx]++
		if shortCheck[idx] == longCheck[idx] {
			match++
		} else if shortCheck[idx] == longCheck[idx]-1 {
			match--
		}
	}
	return match == 26
}

// time complexity: O(n)
// space complexity: O(1)
func checkInclusion(s1 string, s2 string) bool {
	var count [26]int

	for _, val := range s1 {
		count[val-'a']++
	}

	l := 0
	for r := range s2 {
		count[s2[r]-'a']--
		if r+1 > len(s1) {
			count[s2[l]-'a']++
			l++
		}
		if count == [26]int{} {
			return true
		}
	}

	return false
}
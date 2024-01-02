// tags: arrays&hashing, easy

// time complexity: O(n)
// space complexity: O(1)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count [26]int
	for i := range s {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	for idx := range count {
		if count[idx] != 0 {
			return false
		}
	}
	return true
}
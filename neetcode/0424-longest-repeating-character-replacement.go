// time complexity: O(n)
// space complexity: O(n)
func characterReplacement(s string, k int) int {
	count := make(map[string]int)
	var leftIdx, rightIdx int
	var maxF int
	for ; rightIdx < len(s); rightIdx++ {
		count[string(s[rightIdx])]++
		maxF = max(maxF, count[string(s[rightIdx])])

		if rightIdx-leftIdx+1-maxF > k {
			count[string(s[leftIdx])]--
			leftIdx++
		}
	}
	return rightIdx - leftIdx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
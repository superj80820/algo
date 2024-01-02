// tags: backtracking, star1

// time complexity: O(n^2*2^n)
// space complexity: O(n)
func partition(s string) [][]string {
	var res [][]string
	dfs(0, s, []string{}, &res)
	return res
}

func dfs(i int, s string, part []string, res *[][]string) {
	if i == len(s) {
		partForCopy := make([]string, len(part))
		copy(partForCopy, part)
		*res = append(*res, partForCopy)
		return
	}

	for j := i; j < len(s); j++ {
		if isPalindrome(s, i, j) {
			part = append(part, s[i:j+1])
			dfs(j+1, s, part, res)
			part = part[:len(part)-1]
		}
	}
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
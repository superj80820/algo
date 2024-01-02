// tags: backtracking, star1, medium

var numToLetter = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// time complexity: O(n*4^n)
// space complexity: O(n*4^n)
func letterCombinations(digits string) []string {
	var res []string
	dfs(digits, 0, "", &res)
	return res
}

func dfs(digits string, i int, result string, res *[]string) {
	if digits == "" {
		return
	}
	if i == len(digits) {
		*res = append(*res, result)
		return
	}

	for _, letter := range numToLetter[digits[i]] {
		dfs(digits, i+1, result+string(letter), res)
	}
}
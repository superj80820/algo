// time complexity: O(4^n/sqrt(n))
// space complexity: O(4^n/sqrt(n))
// some people say 'time complexity: O(2^n)' and 'space complexity: O(n)'. It's not correct answer. (ref: https://leetcode.com/problems/generate-parentheses/solutions/10100/easy-to-understand-java-backtracking-solution/comments/184522)
// ref:
//  1. https://www.codingninjas.com/studio/library/generate-parentheses
//  2. nth Catalan number
func generateParenthesis(n int) []string {
	res := r(0, 0, n, "", &[]string{})
	return *res
}

func r(leftCount, rightCount, n int, curString string, res *[]string) *[]string {
	if leftCount == n && rightCount == n {
		*res = append(*res, curString)
		return res
	}
	if leftCount < n {
		r(leftCount+1, rightCount, n, curString+"(", res)
	}
	if leftCount > rightCount {
		r(leftCount, rightCount+1, n, curString+")", res)
	}
	return res
}
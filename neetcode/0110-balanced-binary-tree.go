// tags: trees, star2, todo(write), easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time complexity: O(n)
// space complexity: O(n)
func isBalanced(root *TreeNode) bool {
	var dfs func(node *TreeNode) (bool, int)
	dfs = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		leftBalanced, leftDepth := dfs(node.Left)
		rightBalanced, rightDepth := dfs(node.Right)
		balanced := leftBalanced && rightBalanced && math.Abs(float64(leftDepth-rightDepth)) <= 1
		return balanced, max(leftDepth, rightDepth) + 1
	}
	balanced, _ := dfs(root)
	return balanced
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// tags: star3, trees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	res := root.Val

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftVal := dfs(node.Left)
		rightVal := dfs(node.Right)
		res = max(res, node.Val+max(leftVal, 0)+max(rightVal, 0))
		return node.Val + max(leftVal, rightVal, 0)
	}
	dfs(root)

	return res
}

func max(args ...int) int {
	maxVal := args[0]
	for _, arg := range args[1:] {
		if arg > maxVal {
			maxVal = arg
		}
	}
	return maxVal
}
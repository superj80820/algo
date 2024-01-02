// tags: trees, star3, easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var (
		dfs      func(depth int, root *TreeNode)
		maxDepth int
	)
	dfs = func(depth int, root *TreeNode) {
		if root == nil {
			return
		}
		maxDepth = max(maxDepth, depth)
		dfs(depth+1, root.Left)
		dfs(depth+1, root.Right)
	}
	dfs(1, root)

	return maxDepth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
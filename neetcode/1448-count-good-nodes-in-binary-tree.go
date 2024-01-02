// tags: trees, star2, medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time complexity: O(n)
// space complexity: O(d)
// `d` is tree depth
func goodNodes(root *TreeNode) int {
	var count int
	dfs(root, math.MinInt32, &count)
	return count
}

func dfs(root *TreeNode, pathMax int, count *int) {
	if root == nil {
		return
	}
	if root.Val >= pathMax {
		*count++
	}
	pathMax = max(pathMax, root.Val)
	dfs(root.Left, pathMax, count)
	dfs(root.Right, pathMax, count)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
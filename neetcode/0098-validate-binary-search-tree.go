/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time complexity: O(n)
// time complexity: O(d)
// `d` is tree depth
func isValidBST(root *TreeNode) bool {
	return dfs(root, nil, nil)
}

func dfs(node, minVal, maxVal *TreeNode) bool {
	if node == nil {
		return true
	}
	if (minVal != nil && node.Val <= minVal.Val) ||
		(maxVal != nil && node.Val >= maxVal.Val) {
		return false
	}

	return dfs(node.Left, minVal, node) && dfs(node.Right, node, maxVal)
}

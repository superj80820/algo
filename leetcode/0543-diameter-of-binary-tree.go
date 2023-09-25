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
func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	r(root, &res)
	return res
}

func r(root *TreeNode, res *int) int {
	var left, right int
	if root.Left != nil {
		left = r(root.Left, res)
	}
	if root.Right != nil {
		right = r(root.Right, res)
	}

	*res = max(*res, left+right)

	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
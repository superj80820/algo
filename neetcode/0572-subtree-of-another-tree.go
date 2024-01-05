// tags: trees, star2, easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time complexity: O(m * n)
// space complexity: O(m)
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if sameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func sameTree(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	} else if tree1 != nil && tree2 != nil && tree1.Val == tree2.Val {
		return sameTree(tree1.Left, tree2.Left) && sameTree(tree1.Right, tree2.Right)
	}
	return false
}
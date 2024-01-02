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
// space complexity: O(n)
func buildTree(preorder []int, inorder []int) *TreeNode {
	indexMap := make(map[int]int)
	for idx, val := range inorder {
		indexMap[val] = idx
	}

	var build func(offset int, preorder []int, inorder []int) *TreeNode
	build = func(offset int, preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 || len(inorder) == 0 {
			return nil
		}

		mid := indexMap[preorder[0]] - offset
		node := &TreeNode{Val: preorder[0]}
		node.Left = build(offset, preorder[1:mid+1], inorder[:mid])
		node.Right = build(offset+mid+1, preorder[mid+1:], inorder[mid+1:])
		return node
	}

	return build(0, preorder, inorder)
}
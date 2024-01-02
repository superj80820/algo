// tags: trees, star2

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
func kthSmallest(root *TreeNode, k int) int {
	return dfs(root, k, new(int), 0)
}

func dfs(node *TreeNode, k int, count *int, kthSmallestVal int) int {
	if node == nil {
		return kthSmallestVal
	}
	kthSmallestVal = dfs(node.Left, k, count, kthSmallestVal)
	*count++
	if *count == k {
		return node.Val
	}
	kthSmallestVal = dfs(node.Right, k, count, kthSmallestVal)
	return kthSmallestVal
}
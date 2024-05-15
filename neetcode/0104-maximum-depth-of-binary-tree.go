// tags: trees, star3, easy, practice-count:2

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

// time complexity: O(n)
// space complexity: O(n)
func maxDepth(root *TreeNode) int {
	var (
		maxDepth   int
		depthStack []int
		stack      []*TreeNode
	)
	current := root
	curDepth := 1
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			depthStack = append(depthStack, curDepth)
			current = current.Left
			curDepth++
		}
		curDepth = depthStack[len(stack)-1]
		depthStack = depthStack[:len(stack)-1]
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		maxDepth = max(maxDepth, curDepth)

		current = current.Right
		if current != nil {
			curDepth++
		}
	}

	return maxDepth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

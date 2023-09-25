/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time complexity: O(v + e)
// space complexity: O(v)
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		cur := []int{}
		curLen := len(queue)
		for i := 0; i < curLen; i++ {
			head := pop(&queue)
			cur = append(cur, head.Val)
			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}
		res = append(res, cur)
	}

	return res
}

func pop(queue *[]*TreeNode) *TreeNode {
	head := (*queue)[0]
	*queue = (*queue)[1:]
	return head
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time complexity: O(v + e)
// space complexity: O(v)
// func levelOrder(root *TreeNode) [][]int {
// 	res := [][]int{}
// 	if root == nil {
// 		return res
// 	}

// 	dfs(root, -1, &res)

// 	return res
// }

// func dfs(node *TreeNode, depth int, res *[][]int) {
// 	if node == nil {
// 		return
// 	}
// 	curDepth := depth + 1
// 	if len(*res) == curDepth {
// 		*res = append(*res, []int{})
// 	}
// 	(*res)[curDepth] = append((*res)[curDepth], node.Val)
// 	dfs(node.Left, curDepth, res)
// 	dfs(node.Right, curDepth, res)
// }
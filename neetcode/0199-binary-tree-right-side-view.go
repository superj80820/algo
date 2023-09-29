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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		var rightSideVal int
		curQueueLen := len(queue)
		for i := 0; i < curQueueLen; i++ {
			front := deQueue(&queue)
			rightSideVal = front.Val
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
		res = append(res, rightSideVal)

	}
	return res
}

func deQueue(queue *[]*TreeNode) *TreeNode {
	front := (*queue)[0]
	*queue = (*queue)[1:]
	return front
}
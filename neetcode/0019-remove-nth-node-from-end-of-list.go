// tags: linked-list, star1, medium, practice-count:2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity: O(n)
// space complexity: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	startNode, endNode := head, head

	for i := 0; i < n; i++ {
		endNode = endNode.Next
	}

	if endNode == nil {
		return head.Next
	}

	for endNode.Next != nil {
		startNode = startNode.Next
		endNode = endNode.Next
	}

	startNode.Next = startNode.Next.Next

	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time complexity: O(n)
// space complexity: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	startNode, endNode := dummyNode, head

	for i := 1; i < n; i++ {
		endNode = endNode.Next
	}

	for endNode.Next != nil {
		startNode = startNode.Next
		endNode = endNode.Next
	}

	startNode.Next = startNode.Next.Next

	return dummyNode.Next
}
// tags: linked-list, star2, hard

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time complexity: O(n)
// space complexity: O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	groupPre := dummy

	for {
		kth := getKth(groupPre, k)
		if kth == nil {
			break
		}
		groupNext := kth.Next

		var pre, cur *ListNode
		pre, cur = groupNext, groupPre.Next
		for cur != groupNext {
			temp := cur.Next
			cur.Next = pre
			cur, pre = temp, cur
		}

		temp := groupPre.Next
		groupPre.Next = kth
		groupPre = temp
	}

	return dummy.Next
}

func getKth(node *ListNode, k int) *ListNode {
	for node != nil && k > 0 {
		node = node.Next
		k--
	}
	return node
}
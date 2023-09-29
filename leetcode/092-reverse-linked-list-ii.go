/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time complexity: O(n)
// space complexity: O(1)
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	originLeft := dummy
	for i := 1; i < left; i++ {
		originLeft = originLeft.Next
	}

	var pre, cur *ListNode = nil, originLeft.Next
	for i := 0; i <= right-left; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	originLeft.Next.Next = cur
	originLeft.Next = pre

	return dummy.Next
}

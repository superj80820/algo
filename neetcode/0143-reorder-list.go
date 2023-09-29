/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time complexity: O(n)
// space complexity: O(1)
func reorderList(head *ListNode) {
	// find left list start, left list end, right list start
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	leftStart, leftEnd := head, slow
	rightStart := slow.Next

	leftEnd.Next = nil

	// reverse right list
	var pre, cur *ListNode = nil, rightStart
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	rightStart = pre

	// merge left and right list
	for rightStart != nil {
		leftTemp := leftStart.Next
		rightTemp := rightStart.Next

		leftStart.Next = rightStart
		rightStart.Next = leftTemp

		leftStart = leftTemp
		rightStart = rightTemp
	}
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// time complexity: O(n)
// space complexity: O(1)
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	var length int
	var originTail *ListNode
	curNode := head
	for curNode != nil {
		originTail = curNode
		curNode = curNode.Next
		length++
	}
	k = k % length
	if k == 0 {
		return head
	}

	newTail := head
	for i := 1; i < length-k; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	originTail.Next = head

	return newHead
}
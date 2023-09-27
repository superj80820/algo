// /**
//   - Definition for singly-linked list.
//   - type ListNode struct {
//   - Val int
//   - Next *ListNode
//   - }
//     */
//
// time complexity: O(n)
// space complexity: O(1)
// `n` is max length of list node
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		res, preL1, curL1, preL2, curL2 *ListNode
		carry                           int
	)
	curL1 = l1
	curL2 = l2
	for curL1 != nil || curL2 != nil {
		num := carry
		if curL1 != nil {
			num += curL1.Val
		}
		if curL2 != nil {
			num += curL2.Val
		}
		carry = num / 10
		num = num % 10
		if curL1 != nil {
			curL1.Val = num
			preL1 = curL1
			curL1 = curL1.Next
			res = l1
		}
		if curL2 != nil {
			curL2.Val = num
			preL2 = curL2
			curL2 = curL2.Next
			res = l2
		}
	}

	if carry != 0 {
		lastNode := &ListNode{Val: carry}
		if res == l1 {
			preL1.Next = lastNode
		} else {
			preL2.Next = lastNode
		}
	}

	return res
}

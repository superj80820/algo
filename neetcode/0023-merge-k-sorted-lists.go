// tags: linked-list, star3, hard, practice-count:2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity: O(nlogk)
// space complexity: O(k)
// `k` is length of lists
// `n` is length of all node
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		var mergedLists []*ListNode

		for i := 0; i < len(lists); i = i + 2 {
			var list1, list2 *ListNode
			list1 = lists[i]
			if i+1 < len(lists) {
				list2 = lists[i+1]
			}
			mergedLists = append(mergedLists, merge(list1, list2))
		}
		lists = mergedLists
	}

	return lists[0]
}

func merge(list1, list2 *ListNode) *ListNode {
	dummy := new(ListNode)

	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
			cur = cur.Next
		}
	}

	if list1 != nil {
		cur.Next = list1
	} else if list2 != nil {
		cur.Next = list2
	}

	return dummy.Next
}
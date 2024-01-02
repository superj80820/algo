// tags: linked-list, star3, medium

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

// time complexity: O(n)
// space complexity: O(n)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node)

	curNode := head
	for curNode != nil {
		nodeMap[curNode] = &Node{Val: curNode.Val}
		curNode = curNode.Next
	}

	curNode = head
	for curNode != nil {
		copyList := nodeMap[curNode]
		copyList.Next = nodeMap[curNode.Next]
		copyList.Random = nodeMap[curNode.Random]
		curNode = curNode.Next
	}

	return nodeMap[head]
}
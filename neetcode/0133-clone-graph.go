// tags: graphs, star1, medium

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// time complexity: O(e + v)
// space complexity: O(v)
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	return dfs(node, make(map[*Node]*Node))
}

func dfs(node *Node, oldToNew map[*Node]*Node) *Node {
	if newNode, ok := oldToNew[node]; ok {
		return newNode
	}

	newNode := &Node{Val: node.Val}
	oldToNew[node] = newNode

	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, dfs(neighbor, oldToNew))
	}

	return newNode
}
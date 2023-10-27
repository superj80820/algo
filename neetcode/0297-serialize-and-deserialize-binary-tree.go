// tags: trees, star1

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// time complexity: O(n)
// space complexity: O(1)
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var (
		serializeStrings []string
		dfs              func(node *TreeNode)
	)

	dfs = func(node *TreeNode) {
		if node == nil {
			serializeStrings = append(serializeStrings, "n")
			return
		}
		serializeStrings = append(serializeStrings, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	return strings.Join(serializeStrings, ",")
}

// time complexity: O(n)
// space complexity: O(n)
// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	deserializeStrings := strings.Split(data, ",")

	var (
		dfs func() *TreeNode
		i   int
	)
	dfs = func() *TreeNode {
		if deserializeStrings[i] == "n" {
			i++
			return nil
		}
		val, _ := strconv.Atoi(deserializeStrings[i])
		node := TreeNode{Val: val}
		i++
		node.Left = dfs()
		node.Right = dfs()
		return &node
	}

	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
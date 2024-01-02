// tags: graphs, union-find, medium

// time complexity: O(n*a(n)) ~= O(n)
// space complexity: O(n)
// `n` is length of node
// `a` is inverse Ackermann function
func findRedundantConnection(edges [][]int) []int {
	par := make([]int, len(edges)+1)
	rank := make([]int, len(edges)+1)
	for i := 0; i <= len(edges); i++ {
		par[i] = i
		rank[i] = 1
	}

	var find func(n int) int
	find = func(n int) int {
		if n != par[n] {
			par[n] = find(par[n])
		}
		return par[n]
	}
	var union func(a, b int) bool
	union = func(a, b int) bool {
		findA, findB := find(a), find(b)

		if findA == findB {
			return false
		}
		if rank[findA] >= rank[findB] {
			par[findB] = par[findA]
			rank[findA] += rank[findB]
		} else {
			par[findA] = par[findB]
			rank[findB] += rank[findA]
		}

		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}

	return []int{}
}
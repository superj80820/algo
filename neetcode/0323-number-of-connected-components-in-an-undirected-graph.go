// tags: graphs, union-find, leet-code-premium, medium

/**
 * @param n: the number of vertices
 * @param edges: the edges of undirected graph
 * @return: the number of connected components
 */
// time complexity: O(n*a(n)) ~= O(n)
// space complexity: O(n)
// `n` is length of node
// `a` is inverse Ackermann function
func CountComponents(n int, edges [][]int) int {
	par, rank := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = i
		rank[i] = 1
	}
	unionFind := UnionFind{
		par:  par,
		rank: rank,
	}

	for _, edge := range edges {
		unionFind.Union(edge[0], edge[1])
	}

	rootCount := make(map[int]bool)
	for i := 0; i < n; i++ {
		rootCount[unionFind.Find(i)] = true
	}

	return len(rootCount)
}

type UnionFind struct {
	par  []int
	rank []int
}

func (u *UnionFind) Find(n int) int {
	if n != u.par[n] {
		u.par[n] = u.Find(u.par[n])
	}
	return u.par[n]
}

func (u *UnionFind) Union(a, b int) bool {
	findA, findB := u.Find(a), u.Find(b)

	if findA == findB {
		return false
	}
	if u.rank[findA] >= u.rank[findB] {
		u.par[findB] = findA
		u.rank[findA] += u.rank[findB]
	} else {
		u.par[findA] = findB
		u.rank[findB] += u.rank[findA]
	}

	return true
}
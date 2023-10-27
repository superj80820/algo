// tags: graphs, union-find, dfs

/**
 * @param n: An integer
 * @param edges: a list of undirected edges
 * @return: true if it's a valid tree, or false
 */
// time complexity: O(v+e)
// space complexity: O(v+e)
// `v` is length of node
// `e` is length of edges
func ValidTree(n int, edges [][]int) bool {
	adj := make(map[int][]int)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	visited := make(map[int]bool)

	if !dfs(0, -1, adj, visited) {
		return false
	}

	return len(visited) == n
}

func dfs(cur, pre int, adj map[int][]int, visited map[int]bool) bool {
	if _, ok := visited[cur]; ok {
		return false
	}

	visited[cur] = true

	for _, next := range adj[cur] {
		if next != pre && !dfs(next, cur, adj, visited) {
			return false
		}
	}

	return true
}

/**
 * @param n: An integer
 * @param edges: a list of undirected edges
 * @return: true if it's a valid tree, or false
 */
// time complexity: O(n*a(n)) ~= O(n)
// space complexity: O(n)
// `n` is length of node
// `a` is inverse Ackermann function
func ValidTree(n int, edges [][]int) bool {
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
		if !unionFind.Union(edge[0], edge[1]) {
			return false
		}
	}

	root := unionFind.Find(0)
	for i := 1; i < n; i++ {
		if root != unionFind.Find(i) {
			return false
		}
	}

	return true
}

type UnionFind struct {
	par  []int
	rank []int
}

func (u *UnionFind) Find(n int) int {
	if n != u.par[n] {
		u.par[n] = u.Find(u.par[n]) // yorktodo
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
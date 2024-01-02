// tags: heap(priority-queue), star3, dijkstra's-algo, medium

import "container/heap"

type NodeHeap [][2]int                // [][2]int{{weight, node}}
func (n NodeHeap) Len() int           { return len(n) }
func (n NodeHeap) Less(i, j int) bool { return n[i][0] < n[j][0] }
func (n NodeHeap) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n *NodeHeap) Push(x any)        { *n = append(*n, x.([2]int)) }
func (n *NodeHeap) Pop() any {
	x := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return x
}

// time complexity: O(e*logv)
// space complexity: O(e)
func networkDelayTime(times [][]int, n int, k int) int {
	adj := make(map[int][][2]int, n)
	for _, time := range times {
		node, neiNode, weight := time[0], time[1], time[2]
		adj[node] = append(adj[node], [2]int{weight, neiNode})
	}

	var res int
	visited := make(map[int]bool)
	pq := NodeHeap{{0, k}}
	for len(pq) != 0 {
		minNode := heap.Pop(&pq).([2]int)
		weight, node := minNode[0], minNode[1]
		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = true
		res = weight
		for _, nei := range adj[node] {
			neiWeight, neiNode := nei[0], nei[1]
			if _, ok := visited[neiNode]; ok {
				continue
			}
			heap.Push(&pq, [2]int{weight + neiWeight, neiNode})
		}
	}

	if len(visited) != n {
		return -1
	}
	return res
}
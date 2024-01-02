// tags: advanced-graphs, star3, prim's-algo, kruskal-algo(todo), leet-code-premium, heap, medium

import "container/heap"

type PointHeap [][]int

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PointHeap) Push(x any)        { *h = append(*h, x.([]int)) }
func (h *PointHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// time complexity: O(n^2*logn)
// space complexity: O(n)
// `n` is length of points
func minCostConnectPoints(points [][]int) int {
	adj := make([][][2]int, len(points))
	for i := 0; i < len(points); i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < len(points); j++ {
			x2, y2 := points[j][0], points[j][1]
			dist := abs(x1-x2) + abs(y1-y2)
			adj[i] = append(adj[i], [2]int{dist, j})
			adj[j] = append(adj[j], [2]int{dist, i})
		}
	}

	var res int
	visited := make(map[int]bool)
	pq := PointHeap{{0, 0}}
	for len(visited) < len(points) {
		minNode := heap.Pop(&pq).([]int)
		if _, ok := visited[minNode[1]]; ok {
			continue
		}
		res += minNode[0]
		visited[minNode[1]] = true
		for _, nei := range adj[minNode[1]] {
			if _, ok := visited[nei[1]]; ok {
				continue
			}
			heap.Push(&pq, []int{nei[0], nei[1]})
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
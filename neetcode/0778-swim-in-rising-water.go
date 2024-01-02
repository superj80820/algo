// tags: advanced-graphs, star3, dijkstra's-algo, hard

import "container/heap"

type HeapCell []Cell

func (h HeapCell) Len() int           { return len(h) }
func (h HeapCell) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h HeapCell) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *HeapCell) Push(x any) {
	*h = append(*h, x.(Cell))
}
func (h *HeapCell) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type Cell struct {
	Val int
	Row int
	Col int
}

// time complexity: O(n^2logn)
// space complexity: O(n^2)
// `n` is row or col length
func swimInWater(grid [][]int) int {
	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true
	pq := &HeapCell{{Val: grid[0][0]}}
	heap.Init(pq)
	for pq.Len() != 0 {
		minCell := heap.Pop(pq).(Cell)
		for _, direct := range [][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} {
			nextRow, nextCol := minCell.Row+direct[0], minCell.Col+direct[1]
			if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) &&
				!visited[[2]int{nextRow, nextCol}] {
				cellVal := max(minCell.Val, grid[nextRow][nextCol])
				if nextRow == len(grid)-1 && nextCol == len(grid[0])-1 {
					return cellVal
				}
				visited[[2]int{nextRow, nextCol}] = true
				heap.Push(pq, Cell{Val: cellVal, Row: nextRow, Col: nextCol})
			}
		}
	}
	return -1 // not found min time case, it won't run
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

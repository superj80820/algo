// tags: heap(priority-queue), star2

import "container/heap"

type PointInfoHeap []*PointInfo

type PointInfo struct {
	Sqrt  int
	Point []int
}

func (heap PointInfoHeap) Len() int           { return len(heap) }
func (heap PointInfoHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }
func (heap PointInfoHeap) Less(i, j int) bool { return heap[i].Sqrt < heap[j].Sqrt }
func (heap *PointInfoHeap) Push(x any) {
	*heap = append(*heap, x.(*PointInfo))
}
func (heap *PointInfoHeap) Pop() any {
	x := (*heap)[len(*heap)-1]
	(*heap) = (*heap)[:len(*heap)-1]
	return x
}

// time complexity: O(n + k * logn)
// space complexity: O(n)
func kClosest(points [][]int, k int) [][]int {
	sqrtPoints := make([]*PointInfo, len(points))
	for idx, point := range points {
		sqrtPoints[idx] = &PointInfo{
			Sqrt:  point[0]*point[0] + point[1]*point[1],
			Point: point,
		}
	}

	pq := (*PointInfoHeap)(&sqrtPoints)
	heap.Init(pq)

	res := make([][]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(pq).(*PointInfo).Point
	}

	return res
}
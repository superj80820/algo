// tags: intervals, hard

import (
	"container/heap"
	"sort"
)

type IntervalHeap [][]int

func (h IntervalHeap) Len() int           { return len(h) }
func (h IntervalHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntervalHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntervalHeap) Push(x any)        { *h = append(*h, x.([]int)) }
func (h *IntervalHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// time complexity: O(mlogm+nlogn)
// space complexity: O(m+n)
// `m` is length of intervals
// `n` is length of queries
func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	sortQueries := make([]int, len(queries))
	copy(sortQueries, queries)
	sort.Ints(sortQueries)

	var intervalsStartIdx int
	resMap := make(map[int]int)
	pq := new(IntervalHeap)
	for _, query := range sortQueries {
		for intervalsStartIdx < len(intervals) && intervals[intervalsStartIdx][0] <= query {
			l, r := intervals[intervalsStartIdx][0], intervals[intervalsStartIdx][1]
			heap.Push(pq, []int{r - l + 1, r})
			intervalsStartIdx++
		}
		for len(*pq) != 0 && query > (*pq)[0][1] {
			heap.Pop(pq)
		}
		if len(*pq) != 0 {
			resMap[query] = (*pq)[0][0]
		} else {
			resMap[query] = -1
		}
	}

	res := make([]int, len(queries))
	for idx := range queries {
		res[idx] = resMap[queries[idx]]
	}

	return res
}
// tags: heap(priority-queue), star2

import "container/heap"

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// time complexity: O(n * logn)
// space complexity: O(n)
func lastStoneWeight(stones []int) int {
	h := intHeap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		firstMaxVal := heap.Pop(&h)
		secondMaxVal := heap.Pop(&h)
		if firstMaxVal == secondMaxVal {
			continue
		} else {
			heap.Push(&h, abs(firstMaxVal.(int)-secondMaxVal.(int)))
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return h[0]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

import (
	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

// time complexity: O(n * logn)
// space complexity: O(n)
func lastStoneWeight(stones []int) int {
	h := binaryheap.NewWith(func(a, b interface{}) int {
		return -utils.IntComparator(a, b)
	})
	for _, stone := range stones {
		h.Push(stone)
	}
	for h.Size() > 1 {
		firstMaxVal, _ := h.Pop()
		secondMaxVal, _ := h.Pop()
		if firstMaxVal == secondMaxVal {
			continue
		} else {
			h.Push(abs(firstMaxVal.(int) - secondMaxVal.(int)))
		}
	}
	if ans, ok := h.Peek(); !ok {
		return 0
	} else {
		return ans.(int)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
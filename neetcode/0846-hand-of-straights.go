// tags: greedy

import "container/heap"

type HeapInt []int

func (h HeapInt) Len() int           { return len(h) }
func (h HeapInt) Less(i, j int) bool { return h[i] < h[j] }
func (h HeapInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *HeapInt) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *HeapInt) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	count := make(map[int]int)
	for _, val := range hand {
		count[val]++
	}

	var keys []int
	for key := range count {
		keys = append(keys, key)
	}

	pq := (*HeapInt)(&keys)
	heap.Init(pq)
	for pq.Len() != 0 {
		minVal := (*pq)[0]
		for i := minVal; i < minVal+groupSize; i++ {
			if _, ok := count[i]; !ok {
				return false
			}
			count[i]--
			if count[i] == 0 {
				if i != (*pq)[0] {
					return false
				}
				heap.Pop(pq)
			}
		}
	}

	return true
}
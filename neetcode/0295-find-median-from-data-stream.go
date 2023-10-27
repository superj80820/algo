// tags: heap(priority-queue), star2

import "container/heap"

type HeapInts struct {
	lessFunc func(a, b int) bool
	data     []int
}

func (h HeapInts) Len() int           { return len(h.data) }
func (h HeapInts) Less(i, j int) bool { return h.lessFunc(h.data[i], h.data[j]) }
func (h HeapInts) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *HeapInts) Push(x any) {
	h.data = append(h.data, x.(int))
}
func (h *HeapInts) Pop() any {
	x := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return x
}

type MedianFinder struct {
	smallHeap HeapInts
	largeHeap HeapInts
}

func Constructor() MedianFinder {
	return MedianFinder{
		smallHeap: HeapInts{
			lessFunc: func(a, b int) bool {
				return a > b
			},
		},
		largeHeap: HeapInts{
			lessFunc: func(a, b int) bool {
				return a < b
			},
		},
	}
}

// time complexity: O(logn)
// space complexity: O(1)
func (this *MedianFinder) AddNum(num int) {
	if this.largeHeap.Len() > 0 && num > this.largeHeap.data[0] {
		heap.Push(&this.largeHeap, num)
	} else {
		heap.Push(&this.smallHeap, num)
	}
	if this.smallHeap.Len()-this.largeHeap.Len() > 1 {
		smallMax := heap.Pop(&this.smallHeap)
		heap.Push(&this.largeHeap, smallMax)
	}
	if this.largeHeap.Len()-this.smallHeap.Len() > 1 {
		largeMin := heap.Pop(&this.largeHeap)
		heap.Push(&this.smallHeap, largeMin)
	}
}

// time complexity: O(1)
// space complexity: O(1)
func (this *MedianFinder) FindMedian() float64 {
	if this.smallHeap.Len() > this.largeHeap.Len() {
		return float64(this.smallHeap.data[0])
	} else if this.smallHeap.Len() < this.largeHeap.Len() {
		return float64(this.largeHeap.data[0])
	} else {
		return float64(this.smallHeap.data[0]+this.largeHeap.data[0]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
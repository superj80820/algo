import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	heap *IntHeap
	k    int
}

// time complexity: O(n * logn)
// space complexity: O(k)
func Constructor(k int, nums []int) KthLargest {
	intHeap := &IntHeap{}
	heap.Init(intHeap)
	kthLargest := KthLargest{
		heap: intHeap,
		k:    k,
	}
	for _, num := range nums {
		kthLargest.Add(num)
	}
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.heap, val)
	if len(*this.heap) > this.k {
		heap.Pop(this.heap)
	}
	return (*this.heap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
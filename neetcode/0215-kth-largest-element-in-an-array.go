// tags: star1, heap(priority-queue), PR

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// time complexity: O(n+klogn)
// space complexity: O(n)
func findKthLargest(nums []int, k int) int {
	pq := IntHeap(nums)
	heap.Init(&pq)
	var res int
	for i := 0; i < k; i++ {
		res = heap.Pop(&pq).(int)
	}
	return res
}

// time complexity:
//
//	best: O(n)
//	average: O(n)
//	worst: O(n^2)
//
// space complexity: O(1)
func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	left, right := 0, len(nums)-1

	for left < right {
		pivot := partition(nums, left, right)

		if pivot < k {
			left = pivot + 1
		} else if pivot > k {
			right = pivot - 1
		} else {
			break
		}
	}

	return nums[k]
}

func partition(nums []int, left, right int) int {
	pivot, fill := nums[right], left

	for i := left; i < right; i++ {
		if nums[i] <= pivot {
			nums[fill], nums[i] = nums[i], nums[fill]
			fill++
		}
	}

	nums[right], nums[fill] = nums[fill], nums[right]

	return fill
}
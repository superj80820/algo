// tags: arrays&hashing, medium

// time complexity: O(k * n)
// space complexity: O(n)
func topKFrequent(nums []int, k int) []int {
    check := make(map[int]int)
    for _, num := range nums {
        check[num]++
    }

    var res []int
    for i := 0; i < k; i++ {
        var maxKey, maxVal int
        for key, val := range check {
            if val > maxVal {
                maxVal = val
                maxKey = key
            }
        }
        delete(check, maxKey)
        res = append(res, maxKey)
    }

    return res
}

import "sort"

// time complexity: O(nlogn)
// space complexity: O(n)
func topKFrequent(nums []int, k int) []int {
    check := make(map[int]int)
    for _, num := range nums {
        check[num]++
    }

    var orderCheck [][2]int
    for key, val := range check {
        orderCheck = append(orderCheck, [2]int{key, val})
    }
    sort.Slice(orderCheck, func(i, j int) bool {
        return orderCheck[i][1] > orderCheck[j][1]
    })

    var res []int
    for _, val := range orderCheck {
        res = append(res, val[0])
    }

    return res[:k]
}

import "container/heap"

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// time complexity: O(n + k * logn)
// space complexity: O(n)
func topKFrequent(nums []int, k int) []int {
	check := make(map[int]int)
	for _, num := range nums {
		check[num]++
	}

	pq := make(PriorityQueue, len(check))
	i := 0
	for value, priority := range check {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	var res []int
	for i := 0; i < k; i++ {
		first := heap.Pop(&pq).(*Item)
		res = append(res, first.value)
	}

	return res
}

// time complexity: O(n)
// space complexity: O(n)
func topKFrequent(nums []int, k int) []int {
    check := make(map[int]int)
    for _, num := range nums {
        check[num]++
    }

    bucket := make([][]int, len(nums))
    for key, val := range check {
        bucket[val-1] = append(bucket[val-1], key)
    }

    var res []int
    var kCount int
    for i := len(bucket)-1; i >= 0; i-- {
        for _, val := range bucket[i] {
            if kCount >= k {
                break
            }
            res = append(res, val)
            kCount++
        }
    }

    return res
}
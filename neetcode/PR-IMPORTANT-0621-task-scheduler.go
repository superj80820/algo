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

// time complexity: O(n)
// space complexity: O(n)
func leastInterval(tasks []byte, n int) int {
	checkCount := make(map[byte]int)
	for _, task := range tasks {
		checkCount[task]++
	}
	var pq IntHeap
	heap.Init(&pq)
	for _, val := range checkCount {
		heap.Push(&pq, val)
	}

	var (
		time  int
		queue []*FreqWithNextTime
	)
	for pq.Len() != 0 || len(queue) != 0 {
		time++
		if pq.Len() > 0 {
			first := heap.Pop(&pq).(int)
			if first > 1 {
				queue = append(queue, &FreqWithNextTime{
					Freq:     first - 1,
					NextTime: time + n},
				)
			}
		}
		if len(queue) != 0 && time >= queue[0].NextTime {
			first := dequeue(&queue)
			heap.Push(&pq, first.Freq)
		}
	}
	return time
}

func dequeue(queue *[]*FreqWithNextTime) *FreqWithNextTime {
	first := (*queue)[0]
	*queue = (*queue)[1:]
	return first
}

type FreqWithNextTime struct {
	Freq     int
	NextTime int
}

// time complexity: O(n)
// space complexity: O(1)
func leastInterval(tasks []byte, n int) int {
	freqCount := make(map[byte]int)
	for _, task := range tasks {
		freqCount[task]++
	}
	var maxFreq int
	for _, val := range checkCount {
		maxFreq = max(maxFreq, val)
	}
	var maxFreqCount int
	for _, val := range checkCount {
		if val == maxFreq {
			maxFreqCount++
		}
	}
	return max((maxFreq-1)*(n+1)+maxFreqCount, len(tasks))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
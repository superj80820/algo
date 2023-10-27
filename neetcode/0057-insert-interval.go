// tags: intervals

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	for idx, interval := range intervals {
		if newInterval[1] < interval[0] {
			res = append(res, newInterval)
			res = append(res, intervals[idx:]...)
			return res
		} else if newInterval[0] > interval[1] {
			res = append(res, interval)
		} else {
			newInterval = []int{min(newInterval[0], interval[0]), max(newInterval[1], interval[1])}
		}
	}
	res = append(res, newInterval)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
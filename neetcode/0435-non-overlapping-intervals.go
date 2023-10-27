// tags: intervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res int
	preEnd := intervals[0][1]
	for _, interval := range intervals[1:] {
		if interval[0] >= preEnd {
			preEnd = interval[1]
		} else {
			res++
			preEnd = min(preEnd, interval[1])
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
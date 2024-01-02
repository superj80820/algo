// tags: intervals, medium

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		curLast := res[len(res)-1]
		if interval[0] <= curLast[1] {
			res[len(res)-1] = []int{min(curLast[0], interval[0]), max(curLast[1], interval[1])}
		} else {
			res = append(res, interval)
		}
	}
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
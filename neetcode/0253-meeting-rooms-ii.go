// tags: intervals

import "sort"

/**
 * Definition of Interval:
 * type Interval struct {
 *     Start, End int
 * }
 */

/**
 * @param intervals: an array of meeting time intervals
 * @return: the minimum number of conference rooms required
 */
func MinMeetingRooms(intervals []*Interval) int {
	meetingIntervals := make([][2]int, len(intervals)*2)
	for idx, interval := range intervals {
		meetingIntervals[idx*2] = [2]int{interval.Start, 1}
		meetingIntervals[idx*2+1] = [2]int{interval.End, -1}
	}

	sort.Slice(meetingIntervals, func(i, j int) bool {
		if meetingIntervals[i][0] == meetingIntervals[j][0] {
			return meetingIntervals[i][1] < meetingIntervals[j][1]
		}
		return meetingIntervals[i][0] < meetingIntervals[j][0]
	})

	var res, count int
	for _, meetingInterval := range meetingIntervals {
		count += meetingInterval[1]
		res = max(res, count)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
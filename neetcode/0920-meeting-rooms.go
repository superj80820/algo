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
 * @return: if a person could attend all meetings
 */
func CanAttendMeetings(intervals []*Interval) bool {
	if len(intervals) == 0 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	preInterval := intervals[0]
	for _, interval := range intervals[1:] {
		if preInterval.End <= interval.Start {
			preInterval = interval
		} else {
			return false
		}
	}

	return true
}
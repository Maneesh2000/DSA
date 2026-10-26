package intervals

// InsertInterval inserts newInterval into sorted, non-overlapping intervals
// and merges any intervals that overlap it. Touching endpoints are merged.
// The input slice is not modified.
//
// Time: O(n). Space: O(n).
func InsertInterval(intervals []Interval, newInterval Interval) []Interval {
	result := make([]Interval, 0, len(intervals)+1)
	index := 0

	// Intervals entirely before newInterval cannot be affected.
	for index < len(intervals) && intervals[index].End < newInterval.Start {
		result = append(result, intervals[index])
		index++
	}

	// Fold all overlapping intervals into newInterval.
	for index < len(intervals) && intervals[index].Start <= newInterval.End {
		if intervals[index].Start < newInterval.Start {
			newInterval.Start = intervals[index].Start
		}
		if intervals[index].End > newInterval.End {
			newInterval.End = intervals[index].End
		}
		index++
	}

	result = append(result, newInterval)
	return append(result, intervals[index:]...)
}

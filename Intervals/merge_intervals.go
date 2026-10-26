package intervals

// MergeIntervals combines all overlapping intervals and returns them ordered
// by start time. Touching intervals, such as [1, 4] and [4, 5], are merged.
// The input slice is not modified.
//
// Time: O(n log n). Space: O(n).
func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{}
	}

	ordered := sortedCopy(intervals)
	merged := make([]Interval, 0, len(ordered))
	current := ordered[0]

	for _, next := range ordered[1:] {
		if next.Start <= current.End {
			// The ranges overlap, so extend the current range if necessary.
			if next.End > current.End {
				current.End = next.End
			}
			continue
		}

		// A gap starts a new merged range.
		merged = append(merged, current)
		current = next
	}

	return append(merged, current)
}

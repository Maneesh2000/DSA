package intervals

import "sort"

// EraseOverlapIntervals returns the minimum number of intervals that must be
// removed so the remaining intervals do not overlap. An interval ending when
// another starts is allowed, following the usual half-open schedule convention
// (for example, [1, 2) and [2, 3) do not overlap). The input is not modified.
//
// The greedy choice is to keep the interval that ends earliest, leaving the
// most room for future intervals.
// Time: O(n log n). Space: O(n).
func EraseOverlapIntervals(intervals []Interval) int {
	if len(intervals) < 2 {
		return 0
	}

	ordered := append([]Interval(nil), intervals...)
	sort.Slice(ordered, func(first, second int) bool {
		if ordered[first].End == ordered[second].End {
			return ordered[first].Start < ordered[second].Start
		}
		return ordered[first].End < ordered[second].End
	})

	removed := 0
	lastEnd := ordered[0].End
	for _, current := range ordered[1:] {
		if current.Start < lastEnd {
			removed++
			continue
		}

		lastEnd = current.End
	}

	return removed
}

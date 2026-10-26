// Package intervals contains common interview problems involving ranges and
// schedules.
package intervals

import "sort"

// Interval represents a range from Start to End. Callers should provide ranges
// where Start is less than or equal to End. Scheduling functions additionally
// expect meetings with a positive duration (Start < End).
type Interval struct {
	Start int
	End   int
}

// sortedCopy returns intervals ordered by start time without changing input.
func sortedCopy(intervals []Interval) []Interval {
	ordered := append([]Interval(nil), intervals...)
	sort.Slice(ordered, func(first, second int) bool {
		if ordered[first].Start == ordered[second].Start {
			return ordered[first].End < ordered[second].End
		}
		return ordered[first].Start < ordered[second].Start
	})
	return ordered
}

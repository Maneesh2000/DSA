package intervals

import "sort"

// CanAttendMeetings reports whether one person can attend every meeting.
// Meetings that touch at an endpoint do not conflict: [1, 2) and [2, 3) are
// compatible. The input slice is not modified.
//
// Time: O(n log n). Space: O(n).
func CanAttendMeetings(meetings []Interval) bool {
	ordered := sortedCopy(meetings)

	for index := 1; index < len(ordered); index++ {
		if ordered[index].Start < ordered[index-1].End {
			return false
		}
	}

	return true
}

// MinMeetingRooms returns the fewest rooms needed to hold every meeting.
// Meetings ending at the exact time another begins can reuse the same room.
// The input slice is not modified.
//
// Time: O(n log n). Space: O(n).
func MinMeetingRooms(meetings []Interval) int {
	if len(meetings) == 0 {
		return 0
	}

	starts := make([]int, len(meetings))
	ends := make([]int, len(meetings))
	for index, meeting := range meetings {
		starts[index] = meeting.Start
		ends[index] = meeting.End
	}
	sort.Ints(starts)
	sort.Ints(ends)

	rooms, mostRooms, endIndex := 0, 0, 0
	for _, start := range starts {
		if start < ends[endIndex] {
			// This meeting starts before the earliest active meeting ends.
			rooms++
			if rooms > mostRooms {
				mostRooms = rooms
			}
		} else {
			// Reuse the room whose meeting ended first.
			endIndex++
		}
	}

	return mostRooms
}

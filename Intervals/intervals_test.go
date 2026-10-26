package intervals

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	input := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	want := []Interval{{1, 6}, {8, 10}, {15, 18}}
	if got := MergeIntervals(input); !reflect.DeepEqual(got, want) {
		t.Errorf("MergeIntervals() = %v, want %v", got, want)
	}
}

func TestInsertInterval(t *testing.T) {
	input := []Interval{{1, 3}, {6, 9}}
	want := []Interval{{1, 5}, {6, 9}}
	if got := InsertInterval(input, Interval{Start: 2, End: 5}); !reflect.DeepEqual(got, want) {
		t.Errorf("InsertInterval() = %v, want %v", got, want)
	}
}

func TestEraseOverlapIntervals(t *testing.T) {
	input := []Interval{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	if got, want := EraseOverlapIntervals(input), 1; got != want {
		t.Errorf("EraseOverlapIntervals() = %d, want %d", got, want)
	}
}

func TestMeetingRooms(t *testing.T) {
	conflicting := []Interval{{0, 30}, {5, 10}, {15, 20}}
	if CanAttendMeetings(conflicting) {
		t.Error("CanAttendMeetings() = true, want false")
	}
	if got, want := MinMeetingRooms(conflicting), 2; got != want {
		t.Errorf("MinMeetingRooms() = %d, want %d", got, want)
	}

	nonConflicting := []Interval{{7, 10}, {2, 4}}
	if !CanAttendMeetings(nonConflicting) {
		t.Error("CanAttendMeetings() = false, want true")
	}
}

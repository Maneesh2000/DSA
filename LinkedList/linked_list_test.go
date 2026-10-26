package linkedlist

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	got := ToSlice(ReverseList(FromSlice([]int{1, 2, 3})))
	want := []int{3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReverseList() = %v, want %v", got, want)
	}
}

func TestMergeSortedLists(t *testing.T) {
	got := ToSlice(MergeSortedLists(FromSlice([]int{1, 2, 4}), FromSlice([]int{1, 3, 4})))
	want := []int{1, 1, 2, 3, 4, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MergeSortedLists() = %v, want %v", got, want)
	}
}

func TestHasCycle(t *testing.T) {
	head := FromSlice([]int{3, 2, 0, -4})
	if HasCycle(head) {
		t.Error("HasCycle() = true for an acyclic list")
	}

	second := head.Next
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = second
	if !HasCycle(head) {
		t.Error("HasCycle() = false for a cyclic list")
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	got := ToSlice(RemoveNthFromEnd(FromSlice([]int{1, 2, 3, 4, 5}), 2))
	want := []int{1, 2, 3, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveNthFromEnd() = %v, want %v", got, want)
	}

	got = ToSlice(RemoveNthFromEnd(FromSlice([]int{1}), 1))
	if !reflect.DeepEqual(got, []int{}) {
		t.Errorf("RemoveNthFromEnd() = %v, want empty list", got)
	}
}

func TestReorderList(t *testing.T) {
	head := FromSlice([]int{1, 2, 3, 4, 5})
	ReorderList(head)
	want := []int{1, 5, 2, 4, 3}
	if got := ToSlice(head); !reflect.DeepEqual(got, want) {
		t.Errorf("ReorderList() = %v, want %v", got, want)
	}
}

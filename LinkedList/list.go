// Package linkedlist contains common interview problems involving singly
// linked lists.
package linkedlist

// ListNode is one node in a singly linked list of integers.
type ListNode struct {
	Value int
	Next  *ListNode
}

// FromSlice creates a linked list whose values appear in the same order as
// values. It returns nil for an empty slice.
func FromSlice(values []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, value := range values {
		tail.Next = &ListNode{Value: value}
		tail = tail.Next
	}
	return dummy.Next
}

// ToSlice returns the values in an acyclic linked list. Do not use it on a
// list that may contain a cycle.
func ToSlice(head *ListNode) []int {
	values := make([]int, 0)
	for current := head; current != nil; current = current.Next {
		values = append(values, current.Value)
	}
	return values
}

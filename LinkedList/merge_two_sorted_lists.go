package linkedlist

// MergeSortedLists merges two ascending linked lists and returns the merged
// ascending list. It reuses and relinks the input nodes rather than allocating
// a new node for every value.
//
// Time: O(m+n). Space: O(1), excluding the returned list.
func MergeSortedLists(first, second *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for first != nil && second != nil {
		if first.Value <= second.Value {
			tail.Next = first
			first = first.Next
		} else {
			tail.Next = second
			second = second.Next
		}
		tail = tail.Next
	}

	// Only one list can have nodes remaining; attach it in one step.
	if first != nil {
		tail.Next = first
	} else {
		tail.Next = second
	}

	return dummy.Next
}

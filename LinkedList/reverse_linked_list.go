package linkedlist

// ReverseList reverses a singly linked list in place and returns its new head.
//
// previous stores the already-reversed part; current walks through the part
// not yet processed. Before changing current.Next, next saves the remaining
// list so it is not lost.
// Time: O(n). Space: O(1).
func ReverseList(head *ListNode) *ListNode {
	var previous *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}

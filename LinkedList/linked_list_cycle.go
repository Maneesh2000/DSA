package linkedlist

// HasCycle reports whether head belongs to a linked list containing a cycle.
// It uses Floyd's tortoise-and-hare algorithm: in a cyclic list, the pointer
// moving two nodes at a time must eventually meet the one moving one node.
//
// Time: O(n). Space: O(1).
func HasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

package linkedlist

// RemoveNthFromEnd removes the nth node measured from the end of an acyclic
// list and returns the possibly new head. n is one-indexed. If n is not in the
// range [1, list length], the original list is returned unchanged.
//
// The fast pointer is kept n nodes ahead of slow. When fast reaches the end,
// slow is immediately before the node to remove.
// Time: O(n). Space: O(1).
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}

	dummy := &ListNode{Next: head}
	fast := dummy
	for step := 0; step < n; step++ {
		if fast.Next == nil {
			return head
		}
		fast = fast.Next
	}

	slow := dummy
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}

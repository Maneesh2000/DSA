package linkedlist

// ReorderList changes an acyclic list in place from
// L0 -> L1 -> ... -> Ln into L0 -> Ln -> L1 -> Ln-1 -> ... .
//
// Flow:
//  1. Find the end of the first half with slow and fast pointers.
//  2. Reverse the second half.
//  3. Alternate nodes from the first and reversed second halves.
//
// Time: O(n). Space: O(1).
func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// For an odd-length list, slow ends at the middle node.
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := ReverseList(slow.Next)
	slow.Next = nil // Split the list before weaving the two halves together.
	first := head

	for second != nil {
		nextFirst := first.Next
		nextSecond := second.Next

		first.Next = second
		second.Next = nextFirst

		first = nextFirst
		second = nextSecond
	}
}

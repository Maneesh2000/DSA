package stackqueue

// QueueUsingStacks implements a first-in-first-out queue using two stacks.
// It is not safe for concurrent use without external locking.
type QueueUsingStacks struct {
	incoming []int // Newly enqueued values.
	outgoing []int // Values ready to dequeue, with the oldest on top.
}

// NewQueueUsingStacks returns an empty queue.
func NewQueueUsingStacks() *QueueUsingStacks {
	return &QueueUsingStacks{}
}

// Enqueue adds value to the back of the queue.
func (q *QueueUsingStacks) Enqueue(value int) {
	q.incoming = append(q.incoming, value)
}

// Dequeue removes and returns the value at the front. The second result is
// false when the queue is empty.
func (q *QueueUsingStacks) Dequeue() (int, bool) {
	q.moveIncomingToOutgoing()
	if len(q.outgoing) == 0 {
		return 0, false
	}

	last := len(q.outgoing) - 1
	value := q.outgoing[last]
	q.outgoing = q.outgoing[:last]
	return value, true
}

// Peek returns the value at the front without removing it. The second result
// is false when the queue is empty.
func (q *QueueUsingStacks) Peek() (int, bool) {
	q.moveIncomingToOutgoing()
	if len(q.outgoing) == 0 {
		return 0, false
	}
	return q.outgoing[len(q.outgoing)-1], true
}

// Len returns the number of values in the queue.
func (q *QueueUsingStacks) Len() int {
	return len(q.incoming) + len(q.outgoing)
}

// moveIncomingToOutgoing reverses new values only when the front stack is
// empty. Each value moves at most once, making queue operations O(1) amortized.
func (q *QueueUsingStacks) moveIncomingToOutgoing() {
	if len(q.outgoing) != 0 {
		return
	}

	for len(q.incoming) > 0 {
		last := len(q.incoming) - 1
		q.outgoing = append(q.outgoing, q.incoming[last])
		q.incoming = q.incoming[:last]
	}
}

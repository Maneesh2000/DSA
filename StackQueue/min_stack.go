package stackqueue

// MinStack is a stack that can return its current minimum in O(1) time. It is
// not safe for concurrent use without external locking.
type MinStack struct {
	values   []int
	minimums []int // minimums[i] is the minimum of values[0:i+1].
}

// NewMinStack returns an empty minimum-tracking stack.
func NewMinStack() *MinStack {
	return &MinStack{}
}

// Push adds value to the top of the stack.
func (s *MinStack) Push(value int) {
	s.values = append(s.values, value)
	if len(s.minimums) == 0 || value < s.minimums[len(s.minimums)-1] {
		s.minimums = append(s.minimums, value)
	} else {
		s.minimums = append(s.minimums, s.minimums[len(s.minimums)-1])
	}
}

// Pop removes and returns the top value. The second result is false when the
// stack is empty.
func (s *MinStack) Pop() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}

	last := len(s.values) - 1
	value := s.values[last]
	s.values = s.values[:last]
	s.minimums = s.minimums[:last]
	return value, true
}

// Top returns the top value without removing it. The second result is false
// when the stack is empty.
func (s *MinStack) Top() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}
	return s.values[len(s.values)-1], true
}

// Min returns the smallest value currently in the stack. The second result is
// false when the stack is empty.
func (s *MinStack) Min() (int, bool) {
	if len(s.minimums) == 0 {
		return 0, false
	}
	return s.minimums[len(s.minimums)-1], true
}

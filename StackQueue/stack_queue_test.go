package stackqueue

import (
	"reflect"
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	if !IsValidParentheses("{[]}") {
		t.Error("IsValidParentheses() = false, want true")
	}
	if IsValidParentheses("([)]") {
		t.Error("IsValidParentheses() = true, want false")
	}
}

func TestQueueUsingStacks(t *testing.T) {
	queue := NewQueueUsingStacks()
	queue.Enqueue(1)
	queue.Enqueue(2)
	if value, ok := queue.Peek(); !ok || value != 1 {
		t.Fatalf("Peek() = (%d, %t), want (1, true)", value, ok)
	}
	if value, ok := queue.Dequeue(); !ok || value != 1 {
		t.Errorf("Dequeue() = (%d, %t), want (1, true)", value, ok)
	}
	if queue.Len() != 1 {
		t.Errorf("Len() = %d, want 1", queue.Len())
	}
}

func TestMinStack(t *testing.T) {
	stack := NewMinStack()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	if value, ok := stack.Min(); !ok || value != -3 {
		t.Fatalf("Min() = (%d, %t), want (-3, true)", value, ok)
	}
	stack.Pop()
	if value, ok := stack.Top(); !ok || value != 0 {
		t.Errorf("Top() = (%d, %t), want (0, true)", value, ok)
	}
	if value, ok := stack.Min(); !ok || value != -2 {
		t.Errorf("Min() = (%d, %t), want (-2, true)", value, ok)
	}
}

func TestEvaluateRPN(t *testing.T) {
	if value, ok := EvaluateRPN([]string{"2", "1", "+", "3", "*"}); !ok || value != 9 {
		t.Errorf("EvaluateRPN() = (%d, %t), want (9, true)", value, ok)
	}
}

func TestDailyTemperatures(t *testing.T) {
	want := []int{1, 1, 4, 2, 1, 1, 0, 0}
	if got := DailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}); !reflect.DeepEqual(got, want) {
		t.Errorf("DailyTemperatures() = %v, want %v", got, want)
	}
}

func TestDecodeString(t *testing.T) {
	if value, ok := DecodeString("3[a2[c]]"); !ok || value != "accaccacc" {
		t.Errorf("DecodeString() = (%q, %t), want (\"accaccacc\", true)", value, ok)
	}
	if _, ok := DecodeString("3[a"); ok {
		t.Error("DecodeString() accepted an unclosed expression")
	}
}

func TestLargestRectangleArea(t *testing.T) {
	if got, want := LargestRectangleArea([]int{2, 1, 5, 6, 2, 3}), 10; got != want {
		t.Errorf("LargestRectangleArea() = %d, want %d", got, want)
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	want := []int{3, 3, 5, 5, 6, 7}
	if got := MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3); !reflect.DeepEqual(got, want) {
		t.Errorf("MaxSlidingWindow() = %v, want %v", got, want)
	}
}

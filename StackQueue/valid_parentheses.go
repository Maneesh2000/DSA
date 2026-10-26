// Package stackqueue contains common interview problems solved with stacks,
// queues, and monotonic data structures.
package stackqueue

// IsValidParentheses reports whether every opening bracket in text is closed
// by the matching bracket in the correct order. Any non-bracket character
// makes the input invalid.
//
// Time: O(n). Space: O(n).
func IsValidParentheses(text string) bool {
	matchingOpen := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]rune, 0, len(text))

	for _, character := range text {
		switch character {
		case '(', '[', '{':
			stack = append(stack, character)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != matchingOpen[character] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}

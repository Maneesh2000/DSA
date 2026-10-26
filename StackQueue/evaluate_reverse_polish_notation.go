package stackqueue

import "strconv"

// EvaluateRPN evaluates an expression in reverse Polish notation. Tokens may
// be integers or one of +, -, *, /. It returns false for malformed input or
// division by zero. Integer division follows Go's truncation toward zero.
//
// Time: O(n). Space: O(n).
func EvaluateRPN(tokens []string) (int, bool) {
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, false
			}

			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			switch token {
			case "+":
				result = left + right
			case "-":
				result = left - right
			case "*":
				result = left * right
			case "/":
				if right == 0 {
					return 0, false
				}
				result = left / right
			}
			stack = append(stack, result)
		default:
			value, err := strconv.Atoi(token)
			if err != nil {
				return 0, false
			}
			stack = append(stack, value)
		}
	}

	if len(stack) != 1 {
		return 0, false
	}
	return stack[0], true
}

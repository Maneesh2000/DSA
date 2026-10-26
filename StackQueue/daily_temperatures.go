package stackqueue

// DailyTemperatures returns, for each day, how many days must pass until a
// warmer temperature occurs. A day with no future warmer temperature has zero.
//
// The stack stores indices with decreasing temperatures. When a warmer day is
// found, it resolves every cooler day on top of the stack.
// Time: O(n). Space: O(n).
func DailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))

	for day, temperature := range temperatures {
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			previousDay := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[previousDay] = day - previousDay
		}
		stack = append(stack, day)
	}

	return answer
}

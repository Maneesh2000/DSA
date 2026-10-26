package stackqueue

// LargestRectangleArea returns the area of the largest rectangle in a
// histogram with non-negative bar heights.
//
// The monotonic stack stores increasing bar indices. When a shorter bar
// arrives, it determines the right boundary for bars that are taller than it.
// A final zero-height sentinel flushes the remaining bars.
// Time: O(n). Space: O(n).
func LargestRectangleArea(heights []int) int {
	values := append(append([]int(nil), heights...), 0)
	stack := make([]int, 0, len(values))
	bestArea := 0

	for index, height := range values {
		for len(stack) > 0 && height < values[stack[len(stack)-1]] {
			barIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := index
			if len(stack) > 0 {
				width = index - stack[len(stack)-1] - 1
			}

			area := values[barIndex] * width
			if area > bestArea {
				bestArea = area
			}
		}
		stack = append(stack, index)
	}

	return bestArea
}

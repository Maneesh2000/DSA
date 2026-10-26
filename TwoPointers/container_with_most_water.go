package twopointers

// MaxArea returns the largest amount of water held between any two vertical
// lines. heights[i] is the height of the line at index i.
//
// Moving the shorter line is the only move that can possibly improve the area:
// the width always decreases, while the shorter height is the limiting factor.
// Time: O(n). Space: O(1).
func MaxArea(heights []int) int {
	left, right := 0, len(heights)-1
	best := 0

	for left < right {
		height := heights[left]
		if heights[right] < height {
			height = heights[right]
		}

		area := height * (right - left)
		if area > best {
			best = area
		}

		if heights[left] <= heights[right] {
			left++
		} else {
			right--
		}
	}

	return best
}

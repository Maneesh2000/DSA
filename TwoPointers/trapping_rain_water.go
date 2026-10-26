package twopointers

// TrapRainWater returns the total amount of rain water trapped between bars.
// height[i] is the non-negative height of the bar at index i.
//
// The shorter side determines which pointer can be settled: a taller bar on
// the opposite side already exists, so its trapped water is known.
// Time: O(n). Space: O(1).
func TrapRainWater(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] <= height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}

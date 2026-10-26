// Package dynamicprogramming contains common interview problems solved with
// dynamic programming.
package dynamicprogramming

// ClimbStairs returns the number of distinct ways to climb n stairs when each
// move covers one or two stairs. It returns 0 for a negative number of stairs.
// There is one way to climb zero stairs: take no steps.
//
// Time: O(n). Space: O(1).
func ClimbStairs(n int) int {
	if n < 0 {
		return 0
	}

	previous, current := 1, 1 // Ways for 0 and 1 stairs.
	for stair := 2; stair <= n; stair++ {
		previous, current = current, previous+current
	}

	return current
}

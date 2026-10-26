package dynamicprogramming

// HouseRobber returns the largest sum obtainable without selecting adjacent
// values. nums represents non-negative amounts in houses arranged in a line.
//
// At each house, either keep the best previous total or add this house to the
// best total from two houses earlier.
// Time: O(n). Space: O(1).
func HouseRobber(nums []int) int {
	twoBack, oneBack := 0, 0

	for _, amount := range nums {
		robCurrent := twoBack + amount
		skipCurrent := oneBack
		if robCurrent > skipCurrent {
			twoBack, oneBack = oneBack, robCurrent
		} else {
			twoBack, oneBack = oneBack, skipCurrent
		}
	}

	return oneBack
}

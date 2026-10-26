// Package twopointers contains common interview problems solved with the
// two-pointer technique.
package twopointers

// TwoSumSorted finds two different values in a sorted slice whose sum equals
// target. It returns their zero-based indices and true when a pair exists;
// otherwise it returns (0, 0, false). The input slice is not modified.
//
// Time: O(n). Space: O(1).
func TwoSumSorted(numbers []int, target int) (int, int, bool) {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		switch {
		case sum == target:
			return left, right, true
		case sum < target:
			// Increasing the left pointer increases the sum in a sorted slice.
			left++
		default:
			// Decreasing the right pointer decreases the sum in a sorted slice.
			right--
		}
	}

	return 0, 0, false
}

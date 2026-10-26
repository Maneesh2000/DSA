// Package basicprograms contains small, foundational programming exercises.
package basicprograms

// BubbleSort returns values in ascending order using bubble sort. The input
// slice is not modified. After each pass, the largest unsorted value has
// "bubbled" to the end, so the next pass can stop one position earlier.
//
// Time: O(n^2) worst case, O(n) for an already sorted slice. Space: O(n) for
// the returned copy.
func BubbleSort(values []int) []int {
	sorted := append([]int(nil), values...)

	for end := len(sorted) - 1; end > 0; end-- {
		swapped := false
		for index := 0; index < end; index++ {
			if sorted[index] > sorted[index+1] {
				sorted[index], sorted[index+1] = sorted[index+1], sorted[index]
				swapped = true
			}
		}

		// No swaps mean the remaining values are already in order.
		if !swapped {
			break
		}
	}

	return sorted
}

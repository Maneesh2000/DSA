package basicprograms

// InsertionSort returns values in ascending order using insertion sort. The
// input slice is not modified. It grows a sorted left portion by shifting
// larger values right before placing the next value into its correct position.
//
// Time: O(n^2) worst case, O(n) for an already sorted slice. Space: O(n) for
// the returned copy.
func InsertionSort(values []int) []int {
	sorted := append([]int(nil), values...)

	for index := 1; index < len(sorted); index++ {
		value := sorted[index]
		position := index - 1

		for position >= 0 && sorted[position] > value {
			sorted[position+1] = sorted[position]
			position--
		}
		sorted[position+1] = value
	}

	return sorted
}

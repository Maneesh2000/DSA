package basicprograms

// SelectionSort returns values in ascending order using selection sort. The
// input slice is not modified. Each pass finds the smallest value remaining
// and places it at the next sorted position.
//
// Time: O(n^2). Space: O(n) for the returned copy.
func SelectionSort(values []int) []int {
	sorted := append([]int(nil), values...)

	for position := 0; position < len(sorted); position++ {
		smallest := position
		for candidate := position + 1; candidate < len(sorted); candidate++ {
			if sorted[candidate] < sorted[smallest] {
				smallest = candidate
			}
		}

		sorted[position], sorted[smallest] = sorted[smallest], sorted[position]
	}

	return sorted
}

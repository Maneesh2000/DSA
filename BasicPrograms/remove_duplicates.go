package basicprograms

// RemoveDuplicatesSorted returns the unique values from an ascending sorted
// slice, preserving order. The input slice is not modified.
//
// Time: O(n). Space: O(n).
func RemoveDuplicatesSorted(values []int) []int {
	if len(values) == 0 {
		return []int{}
	}

	unique := make([]int, 0, len(values))
	unique = append(unique, values[0])
	for _, value := range values[1:] {
		if value != unique[len(unique)-1] {
			unique = append(unique, value)
		}
	}

	return unique
}

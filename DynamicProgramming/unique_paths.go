package dynamicprogramming

// UniquePaths returns the number of routes from the top-left to the
// bottom-right of a rows-by-columns grid when movement is only right or down.
// It returns 0 when either dimension is not positive.
//
// Time: O(rows*columns). Space: O(columns).
func UniquePaths(rows, columns int) int {
	if rows <= 0 || columns <= 0 {
		return 0
	}

	paths := make([]int, columns)
	for col := range paths {
		paths[col] = 1
	}

	for row := 1; row < rows; row++ {
		for col := 1; col < columns; col++ {
			// paths[col] is from above; paths[col-1] is from the left.
			paths[col] += paths[col-1]
		}
	}

	return paths[columns-1]
}

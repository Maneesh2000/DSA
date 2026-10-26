package dynamicprogramming

// EditDistance returns the fewest single-character insertions, deletions, or
// replacements needed to convert source into target. It is Unicode-safe.
//
// Time: O(m*n). Space: O(m*n).
func EditDistance(source, target string) int {
	sourceRunes := []rune(source)
	targetRunes := []rune(target)
	dp := make([][]int, len(sourceRunes)+1)

	for row := range dp {
		dp[row] = make([]int, len(targetRunes)+1)
		dp[row][0] = row // Delete every source character.
	}
	for col := range dp[0] {
		dp[0][col] = col // Insert every target character.
	}

	for row := 1; row <= len(sourceRunes); row++ {
		for col := 1; col <= len(targetRunes); col++ {
			if sourceRunes[row-1] == targetRunes[col-1] {
				dp[row][col] = dp[row-1][col-1]
				continue
			}

			replace := dp[row-1][col-1]
			delete := dp[row-1][col]
			insert := dp[row][col-1]
			dp[row][col] = min(replace, min(delete, insert)) + 1
		}
	}

	return dp[len(sourceRunes)][len(targetRunes)]
}

// min returns the smaller of two integers.
func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

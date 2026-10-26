package dynamicprogramming

// LongestCommonSubsequence returns the length of the longest sequence shared
// by first and second, without requiring matching characters to be adjacent.
// It handles Unicode input by working with runes.
//
// Time: O(m*n). Space: O(m*n).
func LongestCommonSubsequence(first, second string) int {
	firstRunes := []rune(first)
	secondRunes := []rune(second)
	dp := make([][]int, len(firstRunes)+1)
	for row := range dp {
		dp[row] = make([]int, len(secondRunes)+1)
	}

	for row := 1; row <= len(firstRunes); row++ {
		for col := 1; col <= len(secondRunes); col++ {
			if firstRunes[row-1] == secondRunes[col-1] {
				dp[row][col] = dp[row-1][col-1] + 1
			} else if dp[row-1][col] > dp[row][col-1] {
				dp[row][col] = dp[row-1][col]
			} else {
				dp[row][col] = dp[row][col-1]
			}
		}
	}

	return dp[len(firstRunes)][len(secondRunes)]
}

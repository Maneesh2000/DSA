package dynamicprogramming

// LongestArithmeticSubsequence returns the length of the longest arithmetic
// subsequence in nums. An arithmetic subsequence has the same difference
// between every pair of consecutive selected values. Selected values do not
// need to be adjacent in nums.
//
// dp[i][difference] is the length of the longest arithmetic subsequence that
// ends at index i with the given common difference. For every earlier index j,
// nums[j] can be extended by nums[i]. If no sequence ending at j uses that
// difference, the pair nums[j], nums[i] starts a sequence of length two.
//
// Time: O(n^2). Space: O(n^2) in the worst case.
func LongestArithmeticSubsequence(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	dp := make([]map[int]int, len(nums))
	longest := 2

	for current := range nums {
		dp[current] = make(map[int]int)

		for previous := 0; previous < current; previous++ {
			difference := nums[current] - nums[previous]

			// Every pair is an arithmetic subsequence of length two.
			length := 2
			if previousLength, exists := dp[previous][difference]; exists {
				length = previousLength + 1
			}

			// Multiple previous indices can lead to the same state. Keep the best.
			if length > dp[current][difference] {
				dp[current][difference] = length
			}
			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

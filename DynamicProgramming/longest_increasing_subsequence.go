package dynamicprogramming

// LengthOfLIS returns the length of the longest strictly increasing
// subsequence. The elements do not need to be adjacent in the original slice.
//
// dp[i] stores the best subsequence length ending at index i.
// Time: O(n^2). Space: O(n).
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	longest := 1
	for current := range nums {
		dp[current] = 1
		for previous := 0; previous < current; previous++ {
			if nums[previous] < nums[current] && dp[previous]+1 > dp[current] {
				dp[current] = dp[previous] + 1
			}
		}
		if dp[current] > longest {
			longest = dp[current]
		}
	}

	return longest
}

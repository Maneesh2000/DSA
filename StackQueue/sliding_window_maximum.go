package stackqueue

// MaxSlidingWindow returns the maximum value in every contiguous window of
// length k. It returns an empty slice when k is not within [1, len(nums)].
//
// The deque stores indices whose values are decreasing. The front is always
// the maximum for the current window; smaller values behind a new value can
// never become a future maximum and are discarded.
// Time: O(n). Space: O(k).
func MaxSlidingWindow(nums []int, k int) []int {
	if k <= 0 || k > len(nums) {
		return []int{}
	}

	maximums := make([]int, 0, len(nums)-k+1)
	deque := make([]int, 0, k)

	for right, value := range nums {
		// Remove indices that have fallen out of the current window.
		for len(deque) > 0 && deque[0] <= right-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= value {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, right)

		if right >= k-1 {
			maximums = append(maximums, nums[deque[0]])
		}
	}

	return maximums
}

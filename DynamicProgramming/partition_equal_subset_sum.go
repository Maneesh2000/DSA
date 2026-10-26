package dynamicprogramming

// CanPartitionEqualSubsetSum reports whether nums can be split into two
// subsets with equal totals. Values must be non-negative; a negative value
// returns false because this subset-sum DP uses non-negative targets.
//
// Time: O(n * total sum). Space: O(total sum).
func CanPartitionEqualSubsetSum(nums []int) bool {
	total := 0
	for _, number := range nums {
		if number < 0 {
			return false
		}
		total += number
	}
	if total%2 != 0 {
		return false
	}

	target := total / 2
	possible := make([]bool, target+1)
	possible[0] = true

	for _, number := range nums {
		// Go backward so each number contributes to at most one subset position.
		for sum := target; sum >= number; sum-- {
			possible[sum] = possible[sum] || possible[sum-number]
		}
	}

	return possible[target]
}

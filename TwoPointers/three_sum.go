package twopointers

import "sort"

// ThreeSum returns every unique triplet whose values add up to zero. Each
// triplet is sorted in ascending order, and duplicate triplets are omitted.
// The input slice is not modified.
//
// Time: O(n^2). Space: O(n), for the sorted copy (excluding the result).
func ThreeSum(numbers []int) [][]int {
	values := append([]int(nil), numbers...)
	sort.Ints(values)
	triplets := make([][]int, 0)

	for first := 0; first < len(values)-2; first++ {
		// The same first value would produce triplets already found earlier.
		if first > 0 && values[first] == values[first-1] {
			continue
		}

		left, right := first+1, len(values)-1
		for left < right {
			sum := values[first] + values[left] + values[right]
			switch {
			case sum < 0:
				left++
			case sum > 0:
				right--
			default:
				triplets = append(triplets, []int{values[first], values[left], values[right]})
				left++
				right--

				// Skip repeated second and third values to keep triplets unique.
				for left < right && values[left] == values[left-1] {
					left++
				}
				for left < right && values[right] == values[right+1] {
					right--
				}
			}
		}
	}

	return triplets
}

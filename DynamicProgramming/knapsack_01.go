package dynamicprogramming

// Knapsack01 returns the greatest total value obtainable by selecting items
// whose total weight does not exceed capacity. Each item may be selected at
// most once. It returns 0 for a negative capacity, mismatched input lengths,
// or negative weights.
//
// Iterating capacity backward prevents an item from being selected twice.
// Time: O(number of items * capacity). Space: O(capacity).
func Knapsack01(weights, values []int, capacity int) int {
	if capacity < 0 || len(weights) != len(values) {
		return 0
	}

	bestValue := make([]int, capacity+1)
	for item, weight := range weights {
		if weight < 0 {
			return 0
		}

		for currentCapacity := capacity; currentCapacity >= weight; currentCapacity-- {
			candidate := bestValue[currentCapacity-weight] + values[item]
			if candidate > bestValue[currentCapacity] {
				bestValue[currentCapacity] = candidate
			}
		}
	}

	return bestValue[capacity]
}

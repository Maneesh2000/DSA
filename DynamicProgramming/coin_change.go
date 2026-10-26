package dynamicprogramming

// CoinChange returns the fewest coins needed to make amount, or -1 when it is
// impossible. Coin denominations must be positive; non-positive entries are
// ignored.
//
// Time: O(amount * len(coins)). Space: O(amount).
func CoinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}

	// amount+1 is larger than any possible minimum for a non-negative amount.
	minimumCoins := make([]int, amount+1)
	for value := 1; value <= amount; value++ {
		minimumCoins[value] = amount + 1
	}

	for value := 1; value <= amount; value++ {
		for _, coin := range coins {
			if coin > 0 && coin <= value && minimumCoins[value-coin] != amount+1 {
				candidate := minimumCoins[value-coin] + 1
				if candidate < minimumCoins[value] {
					minimumCoins[value] = candidate
				}
			}
		}
	}

	if minimumCoins[amount] == amount+1 {
		return -1
	}
	return minimumCoins[amount]
}

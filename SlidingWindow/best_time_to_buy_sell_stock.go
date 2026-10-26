// Package slidingwindow contains common interview problems solved by expanding
// and shrinking a contiguous window.
package slidingwindow

// MaxProfit returns the largest profit obtainable by buying once and selling
// once on a later day. It returns 0 when no profitable transaction exists.
//
// minimumPrice represents the best buy day seen before the current sell day.
// Time: O(n). Space: O(1).
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minimumPrice := prices[0]
	bestProfit := 0
	for _, price := range prices[1:] {
		if price-minimumPrice > bestProfit {
			bestProfit = price - minimumPrice
		}
		if price < minimumPrice {
			minimumPrice = price
		}
	}

	return bestProfit
}

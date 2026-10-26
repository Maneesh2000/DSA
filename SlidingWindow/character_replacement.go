package slidingwindow

// CharacterReplacement returns the longest substring that can be made of one
// repeated character by replacing at most k characters. It works with Unicode
// text and returns 0 when k is negative.
//
// A window is valid when every non-majority character can be replaced, meaning
// window length minus the count of its most frequent character is at most k.
// Time: O(n). Space: O(n).
func CharacterReplacement(text string, k int) int {
	if k < 0 {
		return 0
	}

	runes := []rune(text)
	counts := make(map[rune]int)
	left, longest, mostFrequent := 0, 0, 0

	for right, character := range runes {
		counts[character]++
		if counts[character] > mostFrequent {
			mostFrequent = counts[character]
		}

		for right-left+1-mostFrequent > k {
			counts[runes[left]]--
			left++
		}

		if length := right - left + 1; length > longest {
			longest = length
		}
	}

	return longest
}

package slidingwindow

// LengthOfLongestSubstring returns the length of the longest substring with no
// repeated character. It works with Unicode text.
//
// lastSeen maps a character to its latest position. When a duplicate enters
// the window, move the left boundary just after its previous occurrence.
// Time: O(n). Space: O(n).
func LengthOfLongestSubstring(text string) int {
	lastSeen := make(map[rune]int)
	left, longest := 0, 0

	for right, character := range []rune(text) {
		if previous, exists := lastSeen[character]; exists && previous >= left {
			left = previous + 1
		}
		lastSeen[character] = right

		if length := right - left + 1; length > longest {
			longest = length
		}
	}

	return longest
}

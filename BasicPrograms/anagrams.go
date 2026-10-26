package basicprograms

// AreAnagrams reports whether first and second contain exactly the same runes
// with the same counts. Comparison is case-sensitive and includes spaces and
// punctuation.
//
// Time: O(n). Space: O(n).
func AreAnagrams(first, second string) bool {
	firstRunes := []rune(first)
	secondRunes := []rune(second)
	if len(firstRunes) != len(secondRunes) {
		return false
	}

	counts := make(map[rune]int, len(firstRunes))
	for _, character := range firstRunes {
		counts[character]++
	}
	for _, character := range secondRunes {
		counts[character]--
		if counts[character] < 0 {
			return false
		}
	}

	return true
}

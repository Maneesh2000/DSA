package slidingwindow

// ContainsPermutation reports whether text contains a substring that is a
// permutation of pattern. It works with Unicode text. The empty pattern is a
// permutation of every string.
//
// The window is kept at exactly the pattern length, then its rune counts are
// compared to the counts required by pattern.
// Time: O(n + m). Space: O(m).
func ContainsPermutation(pattern, text string) bool {
	patternRunes := []rune(pattern)
	textRunes := []rune(text)
	if len(patternRunes) == 0 {
		return true
	}
	if len(patternRunes) > len(textRunes) {
		return false
	}

	required := make(map[rune]int, len(patternRunes))
	for _, character := range patternRunes {
		required[character]++
	}

	window := make(map[rune]int, len(required))
	matchedKinds := 0
	for right, character := range textRunes {
		window[character]++
		if required[character] > 0 && window[character] == required[character] {
			matchedKinds++
		}

		if right >= len(patternRunes) {
			leaving := textRunes[right-len(patternRunes)]
			if required[leaving] > 0 && window[leaving] == required[leaving] {
				matchedKinds--
			}
			window[leaving]--
		}

		if matchedKinds == len(required) {
			return true
		}
	}

	return false
}

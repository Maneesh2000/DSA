package slidingwindow

// MinimumWindow returns the shortest substring of source containing every rune
// in target with at least its required count. It returns an empty string when
// no such substring exists or target is empty. It works with Unicode text.
//
// The right boundary expands until all required rune counts are present; the
// left boundary then shrinks as far as possible while keeping the window valid.
// Time: O(n + m). Space: O(n + m).
func MinimumWindow(source, target string) string {
	sourceRunes := []rune(source)
	targetRunes := []rune(target)
	if len(targetRunes) == 0 || len(targetRunes) > len(sourceRunes) {
		return ""
	}

	required := make(map[rune]int, len(targetRunes))
	for _, character := range targetRunes {
		required[character]++
	}

	window := make(map[rune]int, len(required))
	requiredKinds, matchedKinds := len(required), 0
	left := 0
	bestStart, bestLength := 0, len(sourceRunes)+1

	for right, character := range sourceRunes {
		window[character]++
		if required[character] > 0 && window[character] == required[character] {
			matchedKinds++
		}

		for matchedKinds == requiredKinds {
			if length := right - left + 1; length < bestLength {
				bestStart, bestLength = left, length
			}

			leaving := sourceRunes[left]
			if required[leaving] > 0 && window[leaving] == required[leaving] {
				matchedKinds--
			}
			window[leaving]--
			left++
		}
	}

	if bestLength == len(sourceRunes)+1 {
		return ""
	}
	return string(sourceRunes[bestStart : bestStart+bestLength])
}

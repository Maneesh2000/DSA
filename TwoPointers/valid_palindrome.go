package twopointers

import "unicode"

// IsPalindrome reports whether text reads the same forward and backward after
// ignoring punctuation, spaces, and letter case. It supports Unicode letters
// and digits.
//
// Time: O(n). Space: O(n), used for Unicode-safe rune indexing.
func IsPalindrome(text string) bool {
	runes := []rune(text)
	left, right := 0, len(runes)-1

	for left < right {
		for left < right && !unicode.IsLetter(runes[left]) && !unicode.IsDigit(runes[left]) {
			left++
		}
		for left < right && !unicode.IsLetter(runes[right]) && !unicode.IsDigit(runes[right]) {
			right--
		}

		if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
			return false
		}
		left++
		right--
	}

	return true
}

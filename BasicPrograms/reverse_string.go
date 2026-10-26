package basicprograms

// ReverseString returns text with its characters in reverse order. It works
// with Unicode input by reversing runes rather than individual UTF-8 bytes.
//
// Time: O(n). Space: O(n).
func ReverseString(text string) string {
	runes := []rune(text)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return string(runes)
}

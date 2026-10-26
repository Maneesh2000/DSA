package stackqueue

import "strings"

// DecodeString expands a bracketed repetition expression such as
// "3[a2[c]]" into "accaccacc". A positive or zero repeat count must appear
// immediately before each opening bracket. It returns false for malformed
// expressions.
//
// Time: O(n + output length). Space: O(n + output length).
func DecodeString(encoded string) (string, bool) {
	countStack := make([]int, 0)
	textStack := make([]string, 0)
	current := ""
	number := 0
	readingNumber := false

	for _, character := range encoded {
		switch {
		case character >= '0' && character <= '9':
			number = number*10 + int(character-'0')
			readingNumber = true
		case character == '[':
			if !readingNumber {
				return "", false
			}
			countStack = append(countStack, number)
			textStack = append(textStack, current)
			current = ""
			number = 0
			readingNumber = false
		case character == ']':
			if readingNumber || len(countStack) == 0 {
				return "", false
			}

			last := len(countStack) - 1
			current = textStack[last] + strings.Repeat(current, countStack[last])
			countStack = countStack[:last]
			textStack = textStack[:last]
		case readingNumber:
			// A number has to describe the repeat count of a following '['.
			return "", false
		default:
			current += string(character)
		}
	}

	if len(countStack) != 0 || readingNumber {
		return "", false
	}
	return current, true
}

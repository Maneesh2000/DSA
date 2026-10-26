package dynamicprogramming

// DecodeWays returns the number of ways to decode a string of digits where
// "1" through "26" map to letters. It returns 0 for an empty string, a string
// containing a non-digit, or an invalid zero-prefixed encoding.
//
// Time: O(n). Space: O(1).
func DecodeWays(digits string) int {
	if len(digits) == 0 || digits[0] == '0' {
		return 0
	}
	for index := range digits {
		if digits[index] < '0' || digits[index] > '9' {
			return 0
		}
	}

	// twoBack and oneBack store the number of ways through i-2 and i-1.
	twoBack, oneBack := 1, 1
	for index := 1; index < len(digits); index++ {
		current := 0
		if digits[index] != '0' {
			current += oneBack
		}

		twoDigitValue := (digits[index-1]-'0')*10 + (digits[index] - '0')
		if twoDigitValue >= 10 && twoDigitValue <= 26 {
			current += twoBack
		}

		twoBack, oneBack = oneBack, current
	}

	return oneBack
}

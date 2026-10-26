package slidingwindow

import "testing"

func TestMaxProfit(t *testing.T) {
	if got, want := MaxProfit([]int{7, 1, 5, 3, 6, 4}), 5; got != want {
		t.Errorf("MaxProfit() = %d, want %d", got, want)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	if got, want := LengthOfLongestSubstring("abcabcbb"), 3; got != want {
		t.Errorf("LengthOfLongestSubstring() = %d, want %d", got, want)
	}
}

func TestCharacterReplacement(t *testing.T) {
	if got, want := CharacterReplacement("AABABBA", 1), 4; got != want {
		t.Errorf("CharacterReplacement() = %d, want %d", got, want)
	}
}

func TestContainsPermutation(t *testing.T) {
	if !ContainsPermutation("ab", "eidbaooo") {
		t.Error("ContainsPermutation() = false, want true")
	}
	if ContainsPermutation("ab", "eidboaoo") {
		t.Error("ContainsPermutation() = true, want false")
	}
}

func TestMinimumWindow(t *testing.T) {
	if got, want := MinimumWindow("ADOBECODEBANC", "ABC"), "BANC"; got != want {
		t.Errorf("MinimumWindow() = %q, want %q", got, want)
	}
	if got := MinimumWindow("a", "aa"); got != "" {
		t.Errorf("MinimumWindow() = %q, want empty string", got)
	}
}

package twopointers

import (
	"reflect"
	"testing"
)

func TestTwoSumSorted(t *testing.T) {
	left, right, found := TwoSumSorted([]int{2, 7, 11, 15}, 9)
	if !found || left != 0 || right != 1 {
		t.Errorf("TwoSumSorted() = (%d, %d, %t), want (0, 1, true)", left, right, found)
	}

	if _, _, found := TwoSumSorted([]int{2, 7, 11, 15}, 10); found {
		t.Error("TwoSumSorted() found a pair when none exists")
	}
}

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("A man, a plan, a canal: Panama") {
		t.Error("IsPalindrome() = false, want true")
	}
	if IsPalindrome("race a car") {
		t.Error("IsPalindrome() = true, want false")
	}
}

func TestMaxArea(t *testing.T) {
	if got, want := MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49; got != want {
		t.Errorf("MaxArea() = %d, want %d", got, want)
	}
}

func TestThreeSum(t *testing.T) {
	got := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	want := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ThreeSum() = %v, want %v", got, want)
	}
}

func TestTrapRainWater(t *testing.T) {
	if got, want := TrapRainWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6; got != want {
		t.Errorf("TrapRainWater() = %d, want %d", got, want)
	}
}

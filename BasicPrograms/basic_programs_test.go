package basicprograms

import (
	"reflect"
	"testing"
)

func TestSortingAlgorithms(t *testing.T) {
	input := []int{5, 1, 4, 2, 8}
	want := []int{1, 2, 4, 5, 8}

	algorithms := map[string]func([]int) []int{
		"bubble":    BubbleSort,
		"selection": SelectionSort,
		"insertion": InsertionSort,
	}
	for name, sortValues := range algorithms {
		t.Run(name, func(t *testing.T) {
			if got := sortValues(input); !reflect.DeepEqual(got, want) {
				t.Errorf("sort result = %v, want %v", got, want)
			}
		})
	}

	if !reflect.DeepEqual(input, []int{5, 1, 4, 2, 8}) {
		t.Errorf("sorting modified input: %v", input)
	}
}

func TestReverseString(t *testing.T) {
	if got, want := ReverseString("Go भाषा"), "ाषाभ oG"; got != want {
		t.Errorf("ReverseString() = %q, want %q", got, want)
	}
}

func TestRemoveDuplicatesSorted(t *testing.T) {
	if got, want := RemoveDuplicatesSorted([]int{1, 1, 2, 2, 3}), []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveDuplicatesSorted() = %v, want %v", got, want)
	}
}

func TestAreAnagrams(t *testing.T) {
	if !AreAnagrams("listen", "silent") {
		t.Error("AreAnagrams() = false, want true")
	}
	if AreAnagrams("rat", "car") {
		t.Error("AreAnagrams() = true, want false")
	}
}

func TestBasicNumbers(t *testing.T) {
	if got, want := Fibonacci(10), 55; got != want {
		t.Errorf("Fibonacci(10) = %d, want %d", got, want)
	}
	if !IsPrime(29) || IsPrime(21) {
		t.Error("IsPrime returned an unexpected result")
	}
}

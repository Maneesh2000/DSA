package dynamicprogramming

import "testing"

func TestClimbStairs(t *testing.T) {
	if got, want := ClimbStairs(5), 8; got != want {
		t.Errorf("ClimbStairs(5) = %d, want %d", got, want)
	}
}

func TestHouseRobber(t *testing.T) {
	if got, want := HouseRobber([]int{2, 7, 9, 3, 1}), 12; got != want {
		t.Errorf("HouseRobber() = %d, want %d", got, want)
	}
}

func TestCoinChange(t *testing.T) {
	if got, want := CoinChange([]int{1, 2, 5}, 11), 3; got != want {
		t.Errorf("CoinChange() = %d, want %d", got, want)
	}
	if got := CoinChange([]int{2}, 3); got != -1 {
		t.Errorf("CoinChange() = %d, want -1", got)
	}
}

func TestLengthOfLIS(t *testing.T) {
	if got, want := LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}), 4; got != want {
		t.Errorf("LengthOfLIS() = %d, want %d", got, want)
	}
}

func TestLongestArithmeticSubsequence(t *testing.T) {
	if got, want := LongestArithmeticSubsequence([]int{9, 4, 7, 2, 10}), 3; got != want {
		t.Errorf("LongestArithmeticSubsequence() = %d, want %d", got, want)
	}
	if got, want := LongestArithmeticSubsequence([]int{3, 6, 9, 12}), 4; got != want {
		t.Errorf("LongestArithmeticSubsequence() = %d, want %d", got, want)
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	if got, want := LongestCommonSubsequence("abcde", "ace"), 3; got != want {
		t.Errorf("LongestCommonSubsequence() = %d, want %d", got, want)
	}
}

func TestKnapsack01(t *testing.T) {
	if got, want := Knapsack01([]int{1, 3, 4, 5}, []int{1, 4, 5, 7}, 7), 9; got != want {
		t.Errorf("Knapsack01() = %d, want %d", got, want)
	}
}

func TestEditDistance(t *testing.T) {
	if got, want := EditDistance("horse", "ros"), 3; got != want {
		t.Errorf("EditDistance() = %d, want %d", got, want)
	}
}

func TestUniquePaths(t *testing.T) {
	if got, want := UniquePaths(3, 7), 28; got != want {
		t.Errorf("UniquePaths() = %d, want %d", got, want)
	}
}

func TestCanPartitionEqualSubsetSum(t *testing.T) {
	if !CanPartitionEqualSubsetSum([]int{1, 5, 11, 5}) {
		t.Error("CanPartitionEqualSubsetSum() = false, want true")
	}
	if CanPartitionEqualSubsetSum([]int{1, 2, 3, 5}) {
		t.Error("CanPartitionEqualSubsetSum() = true, want false")
	}
}

func TestDecodeWays(t *testing.T) {
	if got, want := DecodeWays("226"), 3; got != want {
		t.Errorf("DecodeWays(\"226\") = %d, want %d", got, want)
	}
	if got := DecodeWays("06"); got != 0 {
		t.Errorf("DecodeWays(\"06\") = %d, want 0", got)
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestTraversals(t *testing.T) {
	graph := [][]int{
		{1, 2},
		{0, 3},
		{0, 4},
		{1},
		{2},
	}

	if got, want := BFS(graph, 0), []int{0, 1, 2, 3, 4}; !reflect.DeepEqual(got, want) {
		t.Fatalf("BFS() = %v, want %v", got, want)
	}
	if got, want := DFS(graph, 0), []int{0, 1, 3, 2, 4}; !reflect.DeepEqual(got, want) {
		t.Fatalf("DFS() = %v, want %v", got, want)
	}
}

func TestCycleDetection(t *testing.T) {
	acyclic := [][]int{{1}, {0}, {3}, {2}}
	cyclic := [][]int{{1, 2}, {0, 2}, {0, 1}, {4}, {3}}

	for name, detect := range map[string]func([][]int) bool{
		"BFS": HasCycleBFS,
		"DFS": HasCycleDFS,
	} {
		if detect(acyclic) {
			t.Errorf("%s reported a cycle in an acyclic graph", name)
		}
		if !detect(cyclic) {
			t.Errorf("%s did not report a cycle", name)
		}
	}
}

func TestNumberOfIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0'},
		{'1', '0', '0', '1'},
		{'0', '0', '1', '1'},
	}

	if got, want := NumberOfIslands(grid), 2; got != want {
		t.Fatalf("NumberOfIslands() = %d, want %d", got, want)
	}
}

func TestOrangesRotting(t *testing.T) {
	reachable := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	if got, want := OrangesRotting(reachable), 4; got != want {
		t.Errorf("OrangesRotting(reachable) = %d, want %d", got, want)
	}

	unreachable := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	if got, want := OrangesRotting(unreachable), -1; got != want {
		t.Errorf("OrangesRotting(unreachable) = %d, want %d", got, want)
	}
}

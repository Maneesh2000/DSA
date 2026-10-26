package main

import "fmt"

// HasCycleBFS reports whether an undirected graph contains a cycle.
// The graph is represented as an adjacency list where graph[u] contains the
// neighbours of vertex u. It checks every component, including disconnected
// ones. Each undirected edge must be present in both vertices' lists.
func HasCycleBFS(graph [][]int) bool {
	visited := make([]bool, len(graph))

	for start := range graph {
		if visited[start] {
			continue
		}
		if hasCycleFromBFS(start, graph, visited) {
			return true
		}
	}

	return false
}

// hasCycleFromBFS looks for a cycle in the component that contains start.
func hasCycleFromBFS(start int, graph [][]int, visited []bool) bool {
	type queueItem struct {
		vertex int
		parent int
	}

	queue := []queueItem{{vertex: start, parent: -1}}
	visited[start] = true

	for head := 0; head < len(queue); head++ {
		current := queue[head]

		for _, neighbour := range graph[current.vertex] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, queueItem{vertex: neighbour, parent: current.vertex})
				continue
			}

			// A visited neighbour other than the parent means that a cycle exists.
			if neighbour != current.parent {
				return true
			}
		}
	}

	return false
}

// HasCycleDFS reports whether an undirected graph contains a cycle.
// It uses depth-first search and checks every disconnected component.
func HasCycleDFS(graph [][]int) bool {
	visited := make([]bool, len(graph))

	for start := range graph {
		if !visited[start] && hasCycleFromDFS(start, -1, graph, visited) {
			return true
		}
	}

	return false
}

// hasCycleFromDFS traverses one component. parent prevents the same
// undirected edge from being mistaken for a cycle.
func hasCycleFromDFS(vertex, parent int, graph [][]int, visited []bool) bool {
	visited[vertex] = true

	for _, neighbour := range graph[vertex] {
		if !visited[neighbour] {
			if hasCycleFromDFS(neighbour, vertex, graph, visited) {
				return true
			}
			continue
		}

		if neighbour != parent {
			return true
		}
	}

	return false
}

func main() {
	// This graph has two disconnected components and does not contain a cycle.
	graph := [][]int{
		{1, 2}, // 0
		{0, 3}, // 1
		{0, 4}, // 2
		{1},    // 3
		{2},    // 4
	}

	fmt.Println("Cycle detected (BFS):", HasCycleBFS(graph))
	fmt.Println("Cycle detected (DFS):", HasCycleDFS(graph))
}

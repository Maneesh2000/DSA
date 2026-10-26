package main

// BFS returns the order in which breadth-first search visits vertices starting
// at start. The graph is an adjacency list. It returns nil when start is not a
// valid vertex. Time complexity is O(V + E) for the reachable component.
func BFS(graph [][]int, start int) []int {
	if start < 0 || start >= len(graph) {
		return nil
	}

	visited := make([]bool, len(graph))
	order := make([]int, 0)
	queue := []int{start}
	visited[start] = true

	// head advances through the slice, avoiding an expensive front deletion.
	for head := 0; head < len(queue); head++ {
		vertex := queue[head]
		order = append(order, vertex)

		for _, neighbour := range graph[vertex] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}

	return order
}

// DFS returns the order in which depth-first search visits vertices starting
// at start. Neighbours are visited in the same order they appear in graph.
// It returns nil when start is not a valid vertex.
func DFS(graph [][]int, start int) []int {
	if start < 0 || start >= len(graph) {
		return nil
	}

	visited := make([]bool, len(graph))
	order := make([]int, 0)

	var visit func(vertex int)
	visit = func(vertex int) {
		visited[vertex] = true
		order = append(order, vertex)

		for _, neighbour := range graph[vertex] {
			if !visited[neighbour] {
				visit(neighbour)
			}
		}
	}

	visit(start)
	return order
}

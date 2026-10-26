package main

// OrangesRotting returns the minimum number of minutes needed to rot every
// fresh orange. In grid, 0 is empty, 1 is fresh, and 2 is rotten. It returns
// -1 if a fresh orange cannot be reached. The input grid is not modified.
func OrangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	// Work on a copy so callers can reuse the original grid after this function.
	state := make([][]int, len(grid))
	queue := make([][2]int, 0)
	fresh := 0
	for row := range grid {
		state[row] = append([]int(nil), grid[row]...)
		for col, orange := range state[row] {
			switch orange {
			case 1:
				fresh++
			case 2:
				queue = append(queue, [2]int{row, col})
			}
		}
	}

	minutes := 0
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// Process one breadth-first-search layer per minute.
	for fresh > 0 && len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cell := queue[0]
			queue = queue[1:]

			for _, direction := range directions {
				nextRow := cell[0] + direction[0]
				nextCol := cell[1] + direction[1]
				if nextRow < 0 || nextRow >= len(state) || nextCol < 0 || nextCol >= len(state[nextRow]) {
					continue
				}
				if state[nextRow][nextCol] != 1 {
					continue
				}

				state[nextRow][nextCol] = 2
				fresh--
				queue = append(queue, [2]int{nextRow, nextCol})
			}
		}
		minutes++
	}

	if fresh > 0 {
		return -1
	}
	return minutes
}

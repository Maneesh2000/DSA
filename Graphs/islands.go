package main

// NumberOfIslands counts groups of horizontally or vertically adjacent land
// cells in grid. A land cell is '1' and water is '0'. The function leaves the
// input grid unchanged and returns 0 for an empty grid.
func NumberOfIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	visited := make([][]bool, len(grid))
	for row := range grid {
		visited[row] = make([]bool, len(grid[row]))
	}

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	islands := 0

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != '1' || visited[row][col] {
				continue
			}

			islands++
			visited[row][col] = true
			queue := [][2]int{{row, col}}

			// Mark the complete connected land component before finding another island.
			for head := 0; head < len(queue); head++ {
				cell := queue[head]
				for _, direction := range directions {
					nextRow := cell[0] + direction[0]
					nextCol := cell[1] + direction[1]

					if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[nextRow]) {
						continue
					}
					if grid[nextRow][nextCol] == '1' && !visited[nextRow][nextCol] {
						visited[nextRow][nextCol] = true
						queue = append(queue, [2]int{nextRow, nextCol})
					}
				}
			}
		}
	}

	return islands
}

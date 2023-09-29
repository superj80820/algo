// time complexity: O(m * n)
// space complexity: O(m * n)
func maxAreaOfIsland(grid [][]int) int {
	var maxCount int
	for row, line := range grid {
		for col := range line {
			var count int
			dfs(row, col, grid, &count)
			maxCount = max(maxCount, count)
		}
	}
	return maxCount
}

func dfs(row, col int, grid [][]int, count *int) {
	if grid[row][col] != 1 {
		return
	}
	grid[row][col] = -1
	*count++
	if col-1 >= 0 && grid[row][col-1] == 1 {
		dfs(row, col-1, grid, count)
	}
	if row-1 >= 0 && grid[row-1][col] == 1 {
		dfs(row-1, col, grid, count)
	}
	if col+1 < len(grid[0]) && grid[row][col+1] == 1 {
		dfs(row, col+1, grid, count)
	}
	if row+1 < len(grid) && grid[row+1][col] == 1 {
		dfs(row+1, col, grid, count)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// time complexity: O(m * n)
// space complexity: O(m * n)
func maxAreaOfIsland(grid [][]int) int {
	var maxCount int
	for row, line := range grid {
		for col := range line {
			if grid[row][col] != 1 {
				continue
			}
			var count int
			queue := make([][]int, 0, len(grid)*len(grid[0]))
			queue = append(queue, []int{row, col})
			for len(queue) != 0 {
				front := pop(&queue)
				if grid[front[0]][front[1]] != 1 {
					continue
				}
				grid[front[0]][front[1]] = -1
				count++
				if front[1]-1 >= 0 && grid[front[0]][front[1]-1] == 1 {
					queue = append(queue, []int{front[0], front[1] - 1})
				}
				if front[0]-1 >= 0 && grid[front[0]-1][front[1]] == 1 {
					queue = append(queue, []int{front[0] - 1, front[1]})
				}
				if front[1]+1 < len(grid[0]) && grid[front[0]][front[1]+1] == 1 {
					queue = append(queue, []int{front[0], front[1] + 1})
				}
				if front[0]+1 < len(grid) && grid[front[0]+1][front[1]] == 1 {
					queue = append(queue, []int{front[0] + 1, front[1]})
				}
			}
			maxCount = max(maxCount, count)

		}
	}
	return maxCount
}

func pop(queue *[][]int) []int {
	front := (*queue)[0]
	*queue = (*queue)[1:]
	return front
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
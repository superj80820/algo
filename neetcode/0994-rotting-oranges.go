// time complexity: O(m*n)
// space complexity: O(m*n)
func orangesRotting(grid [][]int) int {
	var (
		rotten       [][2]int
		fresh        int
		isFreshExist bool
	)
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 2 {
				rotten = append(rotten, [2]int{row, col})
			} else if grid[row][col] == 1 {
				isFreshExist = true
				fresh++
			}
		}
	}

	if !isFreshExist {
		return 0
	}

	var count int
	for len(rotten) != 0 {
		curRottenLen := len(rotten)
		for i := 0; i < curRottenLen; i++ {
			first := dequeue(&rotten)
			row, col := first[0], first[1]

			for _, val := range [][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} {
				nextRow, nextCol := row+val[0], col+val[1]
				if (nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0])) &&
					grid[nextRow][nextCol] == 1 {
					grid[nextRow][nextCol] = 2
					rotten = append(rotten, [2]int{nextRow, nextCol})
					fresh--
				}
			}
		}
		count++
	}
	if fresh != 0 {
		return -1
	}
	return count - 1
}

func dequeue(queue *[][2]int) [2]int {
	first := (*queue)[0]
	*queue = (*queue)[1:]
	return first
}
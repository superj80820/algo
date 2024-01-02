// tags: graphs, bfs, leet-code-premium, medium

/**
 * @param rooms: m x n 2D grid
 * @return: nothing
 */

import "math"

// time complexity: O(m*n)
// space complexity: O(m*n)
func WallsAndGates(rooms [][]int) {
	visited := make(map[[2]int]bool)

	var queue [][2]int
	for row := range rooms {
		for col := range rooms[0] {
			if rooms[row][col] == 0 {
				queue = append(queue, [2]int{row, col})
				visited[[2]int{row, col}] = true
			}
		}
	}

	var dist int
	for len(queue) != 0 {
		curLen := len(queue)
		for i := 0; i < curLen; i++ {
			first := dequeue(&queue)

			if rooms[first[0]][first[1]] == math.MaxInt32 {
				rooms[first[0]][first[1]] = dist
			}

			for _, direction := range [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} {
				nextRow, nextCol := first[0]+direction[0], first[1]+direction[1]
				if (nextRow >= 0 && nextRow < len(rooms) && nextCol >= 0 && nextCol < len(rooms[0])) &&
					rooms[nextRow][nextCol] == math.MaxInt32 &&
					!visited[[2]int{nextRow, nextCol}] {
					queue = append(queue, [2]int{nextRow, nextCol})
					visited[[2]int{nextRow, nextCol}] = true
				}
			}
		}
		dist++
	}
}

func dequeue(queue *[][2]int) [2]int {
	first := (*queue)[0]
	*queue = (*queue)[1:]
	return first
}
// time complexity: O(log(m * n))
// space complexity: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1
	for left <= right {
		mid := (right-left)/2 + left
		if matrix[mid][0] > target {
			right = mid - 1
		} else if matrix[mid][0] < target {
			left = mid + 1
		} else {
			return true
		}
	}

	targetRow := right
	if targetRow < 0 {
		return false
	}

	for left, right := 0, len(matrix[targetRow])-1; left <= right; {
		mid := (right-left)/2 + left
		if matrix[targetRow][mid] > target {
			right = mid - 1
		} else if matrix[targetRow][mid] < target {
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}

// time complexity: O(m+n)
// space complexity: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		val := matrix[row][col]
		if val == target {
			return true
		} else if target > val {
			row++
		} else if target < val {
			col--
		}
	}

	return false
}
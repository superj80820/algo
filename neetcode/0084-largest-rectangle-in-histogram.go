// tags: star2, stack

// time complexity: O(n)
// space complexity: O(n)
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	var (
		stack [][2]int // index, height
		res   int
	)
	for idx, height := range heights {
		appendIdx := idx
		if idx != 0 && stack[len(stack)-1][1] > height {
			for len(stack) != 0 && stack[len(stack)-1][1] > height {
				var last [2]int
				stack, last = stack[:len(stack)-1], stack[len(stack)-1]
				appendIdx = last[0]
				res = max(res, (idx-last[0])*last[1])
			}

		}
		stack = append(stack, [2]int{appendIdx, height})
	}

	for len(stack) != 0 {
		var last [2]int
		stack, last = stack[:len(stack)-1], stack[len(stack)-1]
		res = max(res, (len(heights)-last[0])*last[1])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// tags: stack, star1

// time complexity: O(n)
// space complexity: O(n)
// monotonic decreasing stack question
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([][]int, 0, len(temperatures))

	for idx, temperature := range temperatures {
		for len(stack) != 0 && temperature > stack[len(stack)-1][0] {
			peak := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[peak[1]] = idx - peak[1]
		}
		stack = append(stack, []int{temperature, idx})
	}

	return res
}
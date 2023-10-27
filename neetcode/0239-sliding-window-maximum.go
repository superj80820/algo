// tags: star1, sliding-window

// time complexity: O(n)
// space complexity: O(1)
func maxSlidingWindow(nums []int, k int) []int {
	var (
		queue, res []int
		l, r       int
	)
	for r < len(nums) {
		for len(queue) != 0 && nums[r] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, r)

		if l > queue[0] {
			queue = queue[1:]
		}

		if r+1 >= k {
			res = append(res, nums[queue[0]])
			l++
		}

		r++
	}

	return res
}
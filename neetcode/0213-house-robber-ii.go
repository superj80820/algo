// tags: 1d-dp, star3

// time complexity: O(n)
// space complexity: O(1)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	return max(robHelper(nums[0:len(nums)-1]), robHelper(nums[1:len(nums)]))
}

func robHelper(nums []int) int {
	var rob1, rob2 int
	for i := range nums {
		rob1, rob2 = rob2, max(rob1+nums[i], rob2)
	}
	return rob2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
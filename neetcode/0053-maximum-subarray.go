// tags: star2, dp

// time complexity: O(n)
// space complexity: O(1)
func maxSubArray(nums []int) int {
	res := nums[0]
	var cur int
	for _, num := range nums {
		cur += num
		res = max(res, cur)
		if cur <= 0 {
			cur = 0
		}
	}
	return res
}

func max(args ...int) int {
	maxVal := args[0]
	for _, val := range args {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
// tags: 1d-dp, star1

// time complexity: O(n)
// space complexity: O(1)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	preMaxMoney := nums[0]
	for i := 2; i < len(nums); i++ {
		preMaxMoney = max(nums[i-2], preMaxMoney)
		nums[i] = nums[i] + preMaxMoney
	}
	return max(nums[len(nums)-1], nums[len(nums)-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// time complexity: O(n)
// space complexity: O(1)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	var rob1, rob2 int
	for i := range nums {
		rob1, rob2 = rob2, max(nums[i]+rob1, rob2)
	}
	return rob2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// tags: linked-list, star3

// time complexity: O(n)
// space complexity: O(1)
func findDuplicate(nums []int) int {
	slowVal := nums[0]
	fastVal := nums[0]
	for {
		slowVal = nums[slowVal]
		fastVal = nums[nums[fastVal]]
		if slowVal == fastVal {
			break
		}
	}
	fastVal = nums[0]
	for slowVal != fastVal {
		slowVal = nums[slowVal]
		fastVal = nums[fastVal]
	}
	return slowVal
}

// time complexity: O(n)
// space complexity: O(1)
func findDuplicate(nums []int) int {
	for _, value := range nums {
		absValue := abs(value)
		if nums[absValue] < 0 {
			return absValue
		}
		nums[absValue] = nums[absValue] * -1
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
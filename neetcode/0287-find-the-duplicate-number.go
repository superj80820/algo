// tags: linked-list, star3, medium, practice-count:2

// time complexity: O(n)
// space complexity: O(1)
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	fast = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
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
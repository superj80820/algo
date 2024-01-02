// tags: greedy, star3, medium

// time complexity: O(n)
// space complexity: O(1)
func canJump(nums []int) bool {
	l, r := len(nums)-1, len(nums)-1
	for ; l >= 0; l-- {
		distance := r - l
		if nums[l] >= distance {
			r = l
		}
	}
	if r == 0 {
		return true
	}
	return false
}
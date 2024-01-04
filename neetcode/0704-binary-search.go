// tags: binary-search, star3, easy

// time complexity: O(logn)
// space complexity: O(1)
func search(nums []int, target int) int {
	for l, r := 0, len(nums)-1; l <= r; {
		mid := (r + l) / 2
		if target < nums[mid] {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
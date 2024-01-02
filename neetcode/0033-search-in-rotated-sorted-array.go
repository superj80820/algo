// tags: binary-search, star2, medium

// time complexity: O(logn)
// space complexity: O(1)
func search(nums []int, target int) int {
	for leftIdx, rightIdx := 0, len(nums)-1; leftIdx <= rightIdx; {
		midIdx := (leftIdx + rightIdx) / 2

		if nums[midIdx] == target {
			return midIdx
		}

		if nums[leftIdx] <= nums[midIdx] {
			if target < nums[midIdx] && target >= nums[leftIdx] {
				rightIdx = midIdx - 1
			} else {
				leftIdx = midIdx + 1
			}
		} else {
			if target > nums[midIdx] && target <= nums[rightIdx] {
				leftIdx = midIdx + 1
			} else {
				rightIdx = midIdx - 1
			}
		}
	}

	return -1
}
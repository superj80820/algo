// tags: binary-search, star1, medium

// time complexity: O(logn)
// space complexity: O(1)
func findMin(nums []int) int {
	minVal := nums[0]

	for leftIdx, rightIdx := 0, len(nums)-1; leftIdx <= rightIdx; {
		if nums[leftIdx] < nums[rightIdx] {
			minVal = min(minVal, nums[leftIdx])
			break
		}

		midIdx := (leftIdx + rightIdx) / 2
		minVal = min(minVal, nums[midIdx])

		if nums[midIdx] >= nums[leftIdx] {
			leftIdx = midIdx + 1
		} else {
			rightIdx = midIdx - 1
		}
	}

	return minVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
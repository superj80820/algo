// tags: two-pointers

// time complexity: O(n)
// space complexity: O(1)
func maxArea(height []int) int {
	leftIdx, rightIdx := 0, len(height)-1
	maxAreaVal := min(height[leftIdx], height[rightIdx]) * (rightIdx - leftIdx)
	for leftIdx < rightIdx {
		if height[leftIdx] < height[rightIdx] {
			leftIdx++
		} else if height[leftIdx] >= height[rightIdx] {
			rightIdx--
		}
		maxAreaVal = max(maxAreaVal, min(height[leftIdx], height[rightIdx])*(rightIdx-leftIdx))
	}
	return maxAreaVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
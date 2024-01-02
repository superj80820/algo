// tags: sliding-window, star3, hard

// time complexity: O(n)
// space complexity: O(1)
func trap(height []int) int {
	l, r := 0, len(height)-1
	var res, maxLeft, maxRight int
	for l <= r {
		if maxLeft <= maxRight {
			res += max(maxLeft-height[l], 0)
			maxLeft = max(maxLeft, height[l])
			l++
		} else {
			res += max(maxRight-height[r], 0)
			maxRight = max(maxRight, height[r])
			r--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
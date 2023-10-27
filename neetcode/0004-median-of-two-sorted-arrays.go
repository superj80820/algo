// tags: binary-search, star3

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	half := total / 2
	long, short := nums1, nums2
	if len(nums1) < len(nums2) {
		long = nums2
		short = nums1
	}

	l, r := 0, len(short)-1
	for {
		i := (r + l) >> 1
		j := half - i - 2

		shortLeft, longLeft := math.MinInt32, math.MinInt32
		shortRight, longRight := math.MaxInt32, math.MaxInt32
		if i >= 0 {
			shortLeft = short[i]
		}
		if i+1 < len(short) {
			shortRight = short[i+1]
		}
		if j >= 0 {
			longLeft = long[j]
		}
		if j+1 < len(long) {
			longRight = long[j+1]
		}

		if shortLeft <= longRight && longLeft <= shortRight {
			if total%2 == 0 {
				return float64(max(shortLeft, longLeft)+min(shortRight, longRight)) / 2
			} else {
				return float64(min(shortRight, longRight))
			}
		} else if shortLeft > longRight {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
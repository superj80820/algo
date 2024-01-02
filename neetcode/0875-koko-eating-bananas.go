// tags: binary-search, star1

import "math"

// time complexity: O(logn)
// space complexity: O(1)
// `n` is max number of piles
func minEatingSpeed(piles []int, h int) int {
	var minEat int
	maxVal := getMaxVal(piles)
	for left, right := 1, maxVal; left <= right; {
		curEat := (right + left) / 2
		var eatHours int
		for _, pile := range piles {
			eatHours += int(math.Ceil(float64(pile) / float64(curEat)))
		}
		if eatHours <= h {
			minEat = curEat
			right = curEat - 1
		} else if eatHours > h {
			left = curEat + 1
		}
	}
	return minEat
}

func getMaxVal(a []int) int {
	maxVal := a[0]
	for _, val := range a[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
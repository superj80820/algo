func canPartition(nums []int) bool {
	sumNums := sum(nums)
	if sumNums%2 != 0 {
		return false
	}
	target := sum(nums) / 2

	dp := make([]bool, target)
	dp[0] = true
	for _, num := range nums {
		nextDp := make([]bool, target)
		for idx, exist := range dp {
			if !exist {
				continue
			}
			newVal := idx + num
			if newVal == target {
				return true
			}
			nextDp[idx] = true
			if newVal < target {
				nextDp[newVal] = true
			}
		}
		dp = nextDp
	}

	return false
}

func sum(nums []int) int {
	var sumVal int
	for _, num := range nums {
		sumVal += num
	}
	return sumVal
}
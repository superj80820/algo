// time complexity: O(n)
// space complexity: O(n)
func longestConsecutive(nums []int) int {
	check := make(map[int]bool)
	for _, num := range nums {
		check[num] = true
	}

	var maxLen int
	for num, _ := range check {
		if check[num-1] {
			continue
		}
		if !check[num-1] {
			startNum := num
			for check[num+1] {
				num++
			}
			maxLen = max(maxLen, num-startNum+1)
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
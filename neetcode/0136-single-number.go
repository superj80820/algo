// tags: bit-manipulation

// time complexity: O(n)
// space complexity: O(1)
func singleNumber(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = ans ^ nums[i]
	}
	return ans
}

// time complexity: O(n)
// space complexity: O(1)
func singleNumber(nums []int) int {
	var ans int
	for i := 0; i < 32; i++ {
		var count int
		bitMask := 1 << i
		for _, num := range nums {
			if num&bitMask == bitMask {
				count++
			}
		}
		if count == 1 {
			ans += bitMask
		}
	}
	return ans
}

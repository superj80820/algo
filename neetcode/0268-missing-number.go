// tags: bit-manipulation, easy

// time complexity: O(n)
// space complexity: O(1)
func missingNumber(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res = res ^ i ^ nums[i]
	}
	return res
}

// time complexity: O(n)
// space complexity: O(1)
func missingNumber(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res = res + i - nums[i]
	}
	return res
}
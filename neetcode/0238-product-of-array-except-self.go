// tags: arrays&hashing

// time complexity: O(n)
// space complexity: O(n)
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	for idx := range res {
		res[idx] = 1
	}

	prefix, postfix := 1, 1
	for i, j := 0, len(nums)-1; i < len(nums)-1; i, j = i+1, j-1 {
		prefix *= nums[i]
		res[i+1] *= prefix

		postfix *= nums[j]
		res[j-1] *= postfix
	}

	return res
}
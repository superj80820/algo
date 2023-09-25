// time complexity: O(n)
// space complexity: O(1)
func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums) && j < len(nums); {
		if nums[i] != 0 {
			i++
			j = i
		} else {
			for j < len(nums) && nums[j] == 0 {
				j++
			}
			if j < len(nums) {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
	}
}

// time complexity: O(n)
// space complexity: O(1)
func moveZeroes(nums []int) {
	var zeroIdx int
	for notZeroIdx := 0; notZeroIdx < len(nums); notZeroIdx++ {
		if nums[notZeroIdx] != 0 {
			nums[zeroIdx], nums[notZeroIdx] = nums[notZeroIdx], nums[zeroIdx]
			zeroIdx++
		}
	}
}
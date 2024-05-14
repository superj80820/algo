// tags: two-pointers, star3, medium

// time complexity: O(n^2)
// space complexity: O(1)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	for idx := 0; idx < len(nums); idx++ {
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}
		for l, r := idx+1, len(nums)-1; l < r; {
			sum := nums[idx] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[idx], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			}
		}
	}
	return res
}
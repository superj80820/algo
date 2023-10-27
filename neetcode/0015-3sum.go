// tags: sliding-window, star3

// time complexity: O(n)
// space complexity: O(1)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sumNum := nums[i] + nums[l] + nums[r]
			if sumNum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			} else if sumNum < 0 {
				l++
			} else if sumNum > 0 {
				r--
			}
		}
	}
	return res
}
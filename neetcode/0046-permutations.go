// time complexity: O(n * n!)
// space complexity: O(n!)
func permute(nums []int) [][]int {
	var res [][]int

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	for range nums {
		first := dequeue(&nums)

		results := permute(nums)
		for idx := range results {
			results[idx] = append(results[idx], first)
		}
		res = append(res, results...)
		nums = append(nums, first)
	}

	return res
}

func dequeue(slice *[]int) int {
	first := (*slice)[0]
	*slice = (*slice)[1:]
	return first
}
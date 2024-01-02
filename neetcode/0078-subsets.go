// tags: backtracking, star3, medium

// time complexity: O(n*2^n)
// space complexity: O(n*2^n)
func subsets(nums []int) [][]int {
	var res [][]int
	getSubset([]int{}, 0, nums, &res)
	return res
}

func getSubset(result []int, idx int, nums []int, res *[][]int) {
	if idx == len(nums) {
		newResult := make([]int, len(result))
		copy(newResult, result)
		*res = append(*res, newResult)
		return
	}
	result = append(result, nums[idx])
	getSubset(result, idx+1, nums, res)
	result = result[:len(result)-1]
	getSubset(result, idx+1, nums, res)
}
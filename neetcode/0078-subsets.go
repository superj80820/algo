// time complexity: O(2^n)
// time complexity: O(2^n)
func subsets(nums []int) [][]int {
	var res [][]int
	getSubset([]int{}, 0, nums, &res)
	return res
}

func getSubset(result []int, idx int, nums []int, res *[][]int) {
	if idx == len(nums) {
		*res = append(*res, result)
		return
	}
	newResult := make([]int, len(result))
	newRestultWithCurValue := make([]int, len(result))
	copy(newResult, result)
	copy(newRestultWithCurValue, result)
	newRestultWithCurValue = append(newRestultWithCurValue, nums[idx])

	getSubset(newResult, idx+1, nums, res)
	getSubset(newRestultWithCurValue, idx+1, nums, res)
}
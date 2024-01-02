// tags: arrays&hashing, star3, easy

// time complexity: O(n)
// space complexity: O(n)
func twoSum(nums []int, target int) []int {
	check := make(map[int]int)
	for idx, num := range nums {
		if val, ok := check[num]; ok {
			return []int{idx, val}
		} else {
			check[target-num] = idx
		}
	}
	return nil
}
// tags: arrays&hashing, easy

// time complexity: O(n)
// space complexity: O(n)
func containsDuplicate(nums []int) bool {
	check := make(map[int]bool)
	for _, num := range nums {
		if _, ok := check[num]; ok {
			return true
		}
		check[num] = true
	}
	return false
}
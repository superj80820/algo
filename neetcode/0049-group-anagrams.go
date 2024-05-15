// tags: arrays&hashing, medium, practice-count:2

// time complexity: O(m * n)
// space complexity: O(m * n)
// m is `strs` length, n is longest string in `strs`
func groupAnagrams(strs []string) [][]string {
	check := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, val := range str {
			count[val-'a']++
		}
		check[count] = append(check[count], str)
	}

	var res [][]string
	for _, val := range check {
		res = append(res, val)
	}

	return res
}
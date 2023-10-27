// tags: greedy

func partitionLabels(s string) []int {
	check := make(map[rune]int)
	for idx, val := range s {
		check[val] = idx
	}

	var (
		size    int
		lastIdx int
		res     []int
	)
	for idx, val := range s {
		lastIdx = max(lastIdx, check[val])
		size++
		if idx == lastIdx {
			res = append(res, size)
			size = 0
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
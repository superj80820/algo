// tags: greedy, medium

func mergeTriplets(triplets [][]int, target []int) bool {
	check := make(map[int]bool)

	for _, triplet := range triplets {
		if triplet[0] > target[0] || triplet[1] > target[1] || triplet[2] > target[2] {
			continue
		}
		for idx := range triplet {
			if triplet[idx] == target[idx] {
				check[idx] = true
			}
		}
	}

	return len(check) == 3
}
// tags: advanced-graphs, star3, bellman-ford-algo, medium

import "math"

// time complexity: O(k*e)
// space complexity: O(v)
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	weights := make([]int, n)
	for i := 0; i < n; i++ {
		weights[i] = math.MaxInt32
	}
	weights[src] = 0

	for i := 0; i <= k; i++ {
		nextWeights := make([]int, n)
		copy(nextWeights, weights)
		for _, flight := range flights {
			from, to, price := flight[0], flight[1], flight[2]

			if weights[from]+price < nextWeights[to] {
				nextWeights[to] = weights[from] + price
			}
		}
		weights = nextWeights
	}

	if weights[dst] == math.MaxInt32 {
		return -1
	}
	return weights[dst]
}
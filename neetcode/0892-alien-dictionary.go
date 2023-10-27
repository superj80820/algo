// tags: advanced-graphs, star3, PR, YT

import "container/heap"

type HeapBytes []byte

func (h HeapBytes) Len() int           { return len(h) }
func (h HeapBytes) Less(i, j int) bool { return h[i] > h[j] }
func (h HeapBytes) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *HeapBytes) Push(x any)        { *h = append(*h, x.(byte)) }
func (h *HeapBytes) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

/**
 * @param words: a list of words
 * @return: a string which is correct order
 */
func AlienOrder(words []string) string {
	adj := make(map[byte][]byte)
	for _, word := range words {
		for i := range word {
			adj[word[i]] = make([]byte, 0)
		}
	}
	for i := 0; i < len(words)-1; i++ {
		word1, word2 := words[i], words[i+1]
		minLen := min(len(word1), len(word2))
		if len(word1) > len(word2) && word1[:minLen] == word2[:minLen] {
			return ""
		}
		for j := 0; j < minLen; j++ {
			if word1[j] != word2[j] {
				adj[word1[j]] = append(adj[word1[j]], word2[j])
				break
			}
		}
	}

	adjOrder := make(HeapBytes, 0, len(adj))
	for key := range adj {
		adjOrder = append(adjOrder, key)
	}
	heap.Init(&adjOrder)

	visited := make(map[byte]bool)
	var (
		reverseRes []byte
		dfs        func(char byte) bool
	)
	dfs = func(char byte) bool {
		if val, ok := visited[char]; ok {
			return val
		}

		visited[char] = true
		for _, val := range adj[char] {
			if dfs(val) {
				return true
			}
		}
		visited[char] = false
		reverseRes = append(reverseRes, char)

		return false
	}

	for len(adjOrder) != 0 {
		minVal := heap.Pop(&adjOrder).(byte)
		if dfs(minVal) {
			return ""
		}
	}

	var res string
	for i := len(reverseRes) - 1; i >= 0; i-- {
		res += string(reverseRes[i])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
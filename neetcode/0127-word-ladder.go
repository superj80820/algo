// tags: graphs, star2, hard

// time complexity: O(m^2*n)
// space complexity: O(m*n)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	var isExist bool
	for _, word := range wordList {
		if endWord == word {
			isExist = true
		}
	}
	if !isExist {
		return 0
	}

	nei := make(map[string][]string)
	wordList = append(wordList, beginWord)
	for _, word := range wordList {
		for i := range word {
			pattern := word[:i] + "*" + word[i+1:]
			nei[pattern] = append(nei[pattern], word)
		}
	}

	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	res := 1
	for len(queue) != 0 {
		curLen := len(queue)
		for i := 0; i < curLen; i++ {
			var word string
			word, queue = queue[0], queue[1:]
			if word == endWord {
				return res
			}
			for j := range word {
				pattern := word[:j] + "*" + word[j+1:]
				for _, neiWord := range nei[pattern] {
					if _, ok := visited[neiWord]; ok {
						continue
					}
					visited[neiWord] = true
					queue = append(queue, neiWord)
				}
			}
		}
		res++
	}

	return 0
}
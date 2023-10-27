// tags: tries

type WordDictionary struct {
	Child [26]*WordDictionary
	IsEnd bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

// time complexity: O(n)
// space complexity: O(n)
func (this *WordDictionary) AddWord(word string) {
	cur := this
	for _, str := range word {
		key := str - 'a'
		if cur.Child[key] == nil {
			cur.Child[key] = &WordDictionary{}
		}
		cur = cur.Child[key]
	}
	cur.IsEnd = true
}

// time complexity: O(n)
// space complexity: O(n)
func (this *WordDictionary) Search(word string) bool {
	var dfs func(idx int, node *WordDictionary) bool
	dfs = func(idx int, node *WordDictionary) bool {
		for i := idx; i < len(word); i++ {
			if word[i] == '.' {
				for _, child := range node.Child {
					if child == nil {
						continue
					}
					if dfs(i+1, child) {
						return true
					}
				}
				return false
			} else {
				key := word[i] - 'a'
				if node.Child[key] == nil {
					return false
				} else {
					node = node.Child[key]
				}
			}
		}
		return node.IsEnd
	}

	return dfs(0, this)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
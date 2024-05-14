// tags: tries, medium

package neetcode

type Trie struct {
	Child [26]*Trie
	IsEnd bool
}

func Constructor() Trie {
	return Trie{
		Child: [26]*Trie{},
		IsEnd: false,
	}
}

// time complexity: O(n)
// space complexity: O(n)
func (this *Trie) Insert(word string) {
	cur := this
	for _, str := range word {
		key := str - 'a'
		if cur.Child[key] == nil {
			cur.Child[key] = &Trie{}
		}
		cur = cur.Child[key]
	}
	cur.IsEnd = true
}

// time complexity: O(n)
// space complexity: O(1)
func (this *Trie) Search(word string) bool {
	cur := this
	for _, str := range word {
		key := str - 'a'
		if cur.Child[key] == nil {
			return false
		} else {
			cur = cur.Child[key]
		}
	}
	return cur.IsEnd
}

// time complexity: O(n)
// space complexity: O(1)
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, str := range prefix {
		key := str - 'a'
		if cur.Child[key] == nil {
			return false
		} else {
			cur = cur.Child[key]
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

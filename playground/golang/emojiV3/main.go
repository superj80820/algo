package main

import "fmt"

type Trie struct {
	children  map[rune]*Trie
	endOfWord bool
	emoji     string
}

func Constructor() Trie {
	return Trie{
		children:  make(map[rune]*Trie),
		endOfWord: false,
	}
}

func (this *Trie) Insert(words string, target string) {
	cur := this

	for _, word := range words {
		if cur.children[word] == nil {
			childTrie := Constructor()
			cur.children[word] = &childTrie
		}
		cur = cur.children[word]
	}

	cur.endOfWord = true
	cur.emoji = target
}

type TrieFinder struct {
	rootTrie *Trie
	curTrie  *Trie
}

func (this *Trie) CreateFinder() TrieFinder {
	return TrieFinder{
		rootTrie: this,
		curTrie:  this,
	}
}

func (this *TrieFinder) Search(word rune) (bool, string) {
	if this.curTrie.children[word] == nil {
		this.curTrie = this.rootTrie
		return false, ""
	}
	this.curTrie = this.curTrie.children[word]
	if this.curTrie.endOfWord {
		return true, this.curTrie.emoji
	}
	return false, ""
}

func main() {
	t := Constructor()

	t.Insert("U+1F3F3::U+FE0F::U+200D::U+1F308", "🏳️‍🌈")

	// exist case
	words := "「林森北」酒精路跑！大家都喝醉～酒錢這次到底要誰出啦！？U+1F3F3::U+FE0F::U+200D::U+1F308哈哈"
	finder := t.CreateFinder()
	for _, word := range words {
		exist, emoji := finder.Search(word)
		if exist {
			fmt.Println("exist case found: ", emoji)
		}
	}

	// does not exist case
	words = "「林森北」酒精路跑！大家都喝醉～酒錢這次到底要誰出啦！？U+1F3F3::U+FE0F::U+200D::U+1F3xx哈哈"
	finder = t.CreateFinder()
	for _, word := range words {
		exist, emoji := finder.Search(word)
		if exist {
			fmt.Println("does not exist case found: ", emoji)
		}
	}
}

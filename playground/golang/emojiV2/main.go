package main

import (
	"fmt"
)

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

func (this *Trie) Search(words string) (bool, string) {
	cur := this

	for _, word := range words {
		if cur.children[word] == nil {
			return false, ""
		}
		cur = cur.children[word]
		if cur.endOfWord {
			return true, cur.emoji
		}
	}

	return false, ""
}

func main() {
	t := Constructor()

	t.Insert("U+1F3F3::U+FE0F::U+200D::U+1F308", "🏳️‍🌈")

	// Exist case
	words := "「林森北」酒精路跑！大家都喝醉～酒錢這次到底要誰出啦！？U+1F3F3::U+FE0F::U+200D::U+1F308哈哈"
	for i := 0; i < len(words); i++ {
		exist, emoji := t.Search(words[i:])
		if exist {
			fmt.Println("find ", emoji)
		}
	}

	// does not exist case
	words = "「林森北」酒精路跑！大家都喝醉～酒錢這次到底要誰出啦！？U+1F3F3::U+FE0F::U+200D::U+1F3xx哈哈"
	for i := 0; i < len(words); i++ {
		exist, emoji := t.Search(words[i:])
		if exist {
			fmt.Println("find ", emoji)
		}
	}
}

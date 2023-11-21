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
	trie := Constructor()

	trie.Insert("🏳 🌈", "🏳️‍🌈")

	// Exist case
	words := "《中華一番！》每天不間斷 馬拉松直播🏳 🌈哈哈"
	for i := 0; i < len(words); i++ {
		exist, emoji := trie.Search(words[i:])
		if exist {
			fmt.Println("find ", emoji)
		}
	}

	// does not exist case
	words = "《中華一番！》每天不間斷 馬拉松直播🏳 哈哈"
	for i := 0; i < len(words); i++ {
		exist, emoji := trie.Search(words[i:])
		if exist {
			fmt.Println("find ", emoji)
		}
	}
}

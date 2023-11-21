package main

import (
	"container/list"
	"fmt"
	"strings"
)

type Trie struct {
	children  map[string]*Trie
	endOfWord bool
	emoji     string
}

func Constructor() Trie {
	return Trie{
		children:  make(map[string]*Trie),
		endOfWord: false,
	}
}

func (this *Trie) Insert(words string, target string) {
	cur := this

	for _, word := range words {
		if cur.children[string(word)] == nil {
			childTrie := Constructor()
			cur.children[string(word)] = &childTrie
		}
		cur = cur.children[string(word)]
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

func (this *TrieFinder) Search(word string) (exist bool, isEnd bool, emoji string) {
	if this.curTrie.children[word] == nil {
		this.curTrie = this.rootTrie
		return false, false, ""
	}
	this.curTrie = this.curTrie.children[word]
	if this.curTrie.endOfWord {
		return true, true, this.curTrie.emoji
	}
	return true, false, ""
}

func Convert(words string, trie *Trie) string {
	formatStringBuilder := list.New()
	finder := trie.CreateFinder()

	runeWords := []rune(words)
	for i := 0; i < len(runeWords); i++ {
		exist, _, _ := finder.Search(string(runeWords[i]))
		if !exist || i == len(runeWords)-1 {
			formatStringBuilder.PushBack(string(runeWords[i]))
		} else {
			emojiBuilder := list.New()
			emojiBuilder.PushBack(string(runeWords[i]))
			for i++; i < len(runeWords); i++ {
				exist, isEnd, emoji := finder.Search(string(runeWords[i]))
				if !exist {
					formatStringBuilder.PushBackList(emojiBuilder)
					break
				} else if isEnd {
					formatStringBuilder.PushBack(emoji)
					break
				}
				emojiBuilder.PushBack(string(runeWords[i]))
			}
		}
	}

	var strBuilder strings.Builder
	for e := formatStringBuilder.Front(); e != nil; e = e.Next() {
		strBuilder.WriteString(e.Value.(string))
	}

	return strBuilder.String()
}

func main() {
	trie := Constructor()

	trie.Insert("🏳 🌈", "🏳️‍🌈")
	trie.Insert("👨‍👩👧", "👨‍👩‍👧")

	// exist case
	fmt.Println(Convert("《中華一番！》每天不間斷👨‍👩👧 馬拉松直播🏳 🌈哈哈", &trie))
	// print: 《中華一番！》每天不間斷👨‍👩‍👧 馬拉松直播🏳️‍🌈哈哈

	// does not exist case
	fmt.Println(Convert("《中華一番！》每天不間斷 馬拉松直播👧 哈哈🏳", &trie))
	// print:  中華一番！》每天不間斷 馬拉松直播👧 哈哈🏳
}

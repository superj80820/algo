package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

type Trie struct {
	children  map[uint64]*Trie
	endOfWord bool
}

func Constructor() Trie {
	return Trie{
		children:  make(map[uint64]*Trie),
		endOfWord: false,
	}
}

func (this *Trie) Insert(words []string) error {
	cur := this

	unicodes, err := this.covertAllStringToInt(words)
	if err != nil {
		return errors.Wrap(err, "error unicode format")
	}

	for _, unicode := range unicodes {
		if cur.children[unicode] == nil {
			childTrie := Constructor()
			cur.children[unicode] = &childTrie
		}
		cur = cur.children[unicode]
	}

	cur.endOfWord = true

	return nil
}

func (this *Trie) Search(words []string) (bool, error) {
	cur := this

	unicodes, err := this.covertAllStringToInt(words)
	if err != nil {
		return false, errors.Wrap(err, "error unicode format")
	}

	for _, unicode := range unicodes {
		if cur.children[unicode] == nil {
			return false, nil
		}
		cur = cur.children[unicode]
	}

	return cur.endOfWord, nil
}

func (this *Trie) StartsWith(words []string) (bool, error) {
	cur := this

	unicodes, err := this.covertAllStringToInt(words)
	if err != nil {
		return false, errors.Wrap(err, "error unicode format")
	}

	for _, unicode := range unicodes {
		if cur.children[unicode] == nil {
			return false, nil
		}
		cur = cur.children[unicode]
	}

	return true, nil
}

func (this *Trie) covertAllStringToInt(words []string) ([]uint64, error) {
	var unicodes []uint64
	for _, word := range words {
		unicode, err := covertStringToInt(word)
		if err != nil {
			return nil, errors.Wrap(err, "error unicode format")
		}
		unicodes = append(unicodes, unicode)
	}
	return unicodes, nil
}

func covertStringToInt(str string) (uint64, error) {
	if str[:2] != "U+" {
		return 0, errors.New("error unicode format")
	}

	unicodeInt, err := strconv.ParseUint(str[2:], 16, 32)
	if err != nil {
		return 0, errors.Wrap(err, "covert failed")
	}
	return unicodeInt, nil
}

func main() {
	t := Constructor()

	t.Insert([]string{"U+1F3F3", "U+FE0F", "U+200D", "U+1F308"})

	fmt.Println(t.Search([]string{"U+1F3F3", "U+FE0F", "U+200D", "U+1F308"}))
	fmt.Println(t.Search([]string{"U+1F3F3", "U+FE0F", "U+200D1", "U+1F308"}))

	fmt.Println(t.StartsWith([]string{"U+1F3F3", "U+FE0F", "U+200D"}))
	fmt.Println(t.StartsWith([]string{"U+1F3F3", "U+FE0F", "U+200Da"}))
}

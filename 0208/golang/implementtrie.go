package implementtrie

import (
	"regexp"
)

type Trie struct {
	wordDict map[string]bool
	wordList []string
	a        int
}

const a int = int('a')

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		wordDict: map[string]bool{},
		wordList: make([]string, 26),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if _, found := this.wordDict[word]; found {
		return
	}

	this.wordDict[word] = true

	idx := int(word[0]) - a
	this.wordList[idx] = this.wordList[idx] + word + ";"
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	_, found := this.wordDict[word]
	return found
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	idx := int(prefix[0]) - a
	charString := this.wordList[idx]

	if len(charString) < len(prefix) {
		return false
	}

	if charString[0:len(prefix)] == prefix {
		return true
	}

	pattern := ";" + prefix
	match, _ := regexp.MatchString(pattern, charString)
	return match
}

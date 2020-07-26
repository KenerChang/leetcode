package addandsearchword

import (
// "fmt"
)

const a = 'a'
const dot = '.'

type WordDictionary struct {
	Nexts        []*WordDictionary
	HasTerminate bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		Nexts: make([]*WordDictionary, 26),
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	var nowChar *WordDictionary = this
	var next *WordDictionary
	for i, char := range word {
		charIdx := char - a
		next = nowChar.Nexts[charIdx]
		if next == nil {
			wordDict := Constructor()
			nowChar.Nexts[charIdx] = &wordDict
			next = &wordDict
		}
		nowChar = next
		if i == len(word)-1 {
			nowChar.HasTerminate = true
		}
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	return this.SearchWord(word)
}

func (this *WordDictionary) SearchWord(word string) bool {
	// fmt.Printf("word: %s\n", word)
	if word == "" {
		return this.HasTerminate
	}

	target := word[0]
	if target == dot {
		for _, next := range this.Nexts {
			if next == nil {
				continue
			}
			found := next.SearchWord(word[1:])
			if found {
				return true
			}
		}
		return false
	}

	charIdx := target - a

	if this.Nexts[charIdx] == nil {
		return false
	}

	next := this.Nexts[charIdx]
	return next.SearchWord(word[1:])
}

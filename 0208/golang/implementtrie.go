package implementtrie

type Trie struct {
	children    [26]*Trie
	isEndOfWord bool
	root        *Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	r := Trie{children: [26]*Trie{}, isEndOfWord: false, root: &Trie{}}
	return r
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	t := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if t.children[index] == nil {
			t.children[index] = &Trie{}
		}
		t = t.children[index]
	}
	t.isEndOfWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	t := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if t.children[index] == nil {
			return false
		}
		t = t.children[index]
	}
	if t.isEndOfWord == true {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	t := this.root
	if len(prefix) == 1 {
		if t.children[prefix[0]-'a'] != nil {
			return true
		}
	}
	depth := 0
	for i := 0; i < len(prefix); i++ {
		depth++
		index := prefix[i] - 'a'
		if t.children[index] == nil {
			return false
		}
		if depth == len(prefix) {
			return true
		}
		t = t.children[index]
	}
	return false
}

package main

import "fmt"

func main() {
	root := Constructor()
	root.Insert("apple")
	fmt.Println(root.Search("app"))
	fmt.Println(root.StartsWith("app"))
}

type Trie struct {
	data  byte
	child []*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		data:  '/',
		child: make([]*Trie, 26, 26),
		isEnd: false,
	}
}

func newTrie(data byte) *Trie {
	return &Trie{
		data:  data,
		child: make([]*Trie, 26, 26),
	}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	p := t
	for _, c := range word {
		index := c - 'a'
		if p.child[index] == nil {
			p.child[index] = newTrie(byte(c))
		}
		p = p.child[index]
	}
	p.isEnd = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	p := t
	for _, c := range word {
		index := c - 'a'
		if p.child[index] == nil {
			return false
		}
		p = p.child[index]
	}
	if !p.isEnd {
		return false
	} else {
		return true
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	p := t
	for _, c := range prefix {
		index := c - 'a'
		if p.child[index] == nil {
			return false
		}
		p = p.child[index]
	}
	return true
}

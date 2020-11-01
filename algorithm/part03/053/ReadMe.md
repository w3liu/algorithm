## Algorithm
### 1. 题目
```
208. 实现 Trie (前缀树)
```
### 2. 题目描述
```
实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");   
trie.search("app");     // 返回 true
说明:

你可以假设所有的输入都是由小写字母 a-z 构成的。
保证所有输入均为非空字符串。
```

### 3. 解答：
```golang

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

```
### 4. 说明
1. 定义Trie的结构中包含的child
2. child数组用26个英文字符表示，用字符减去`a`就得到数组中的索引
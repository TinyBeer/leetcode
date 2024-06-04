package main

/*
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
*/

type Trie struct {
	prefix   byte
	end      bool
	children map[byte]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.end = true
		return
	}
	b := word[0]
	if len(this.children) == 0 {
		this.children = make(map[byte]*Trie)
	}
	child, ok := this.children[b]
	if !ok {
		child = &Trie{prefix: b}
		this.children[b] = child
	}

	child.Insert(word[1:])
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.end
	}
	if len(this.children) == 0 {
		return false
	}
	b := word[0]
	child, ok := this.children[b]
	if !ok {
		return false
	}
	return child.Search(word[1:])
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	if len(this.children) == 0 {
		return false
	}
	b := prefix[0]
	child, ok := this.children[b]
	if !ok {
		return false
	}
	return child.StartsWith(prefix[1:])
}

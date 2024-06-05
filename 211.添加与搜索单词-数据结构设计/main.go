package main

/*
请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。

实现词典类 WordDictionary ：

WordDictionary() 初始化词典对象
void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回
false 。word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。
*/

type WordDictionary struct {
	children [26]*Node
}

type Node struct {
	b        byte
	end      bool
	children [26]*Node
}

func (n *Node) add(word string) {
	if len(word) == 0 {
		n.end = true
		return
	}
	b := word[0]
	idx := int(b - 'a')
	if n.children[idx] == nil {
		n.children[idx] = &Node{b: b}
	}
	n.children[idx].add(word[1:])
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	b := word[0]
	idx := int(b - 'a')
	if this.children[idx] == nil {
		this.children[idx] = &Node{b: b}
	}
	this.children[idx].add(word[1:])
}

func (n *Node) search(word string) bool {
	if len(word) == 0 {
		return n.end
	}
	b := word[0]
	if b == '.' {
		for _, c := range n.children {
			if c != nil && c.search(word[1:]) {
				return true
			}
		}
	} else {
		idx := int(b - 'a')
		return n.children[idx] != nil && n.children[idx].search(word[1:])
	}
	return false
}

func (this *WordDictionary) Search(word string) bool {
	b := word[0]
	if b == '.' {
		for _, c := range this.children {
			if c != nil && c.search(word[1:]) {
				return true
			}
		}
	} else {
		idx := int(b - 'a')
		return this.children[idx] != nil && this.children[idx].search(word[1:])
	}
	return false
}
